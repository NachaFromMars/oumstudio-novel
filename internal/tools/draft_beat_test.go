package tools

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/voocel/ainovel-cli/internal/store"
)

// beatArgs dựng args hợp lệ mặc định cho draft_beat; test chỉnh từng trường để tạo case lỗi.
func beatArgs(t *testing.T, mutate func(m map[string]any)) json.RawMessage {
	t.Helper()
	longVi := strings.Repeat("Gió đêm lùa qua khe cửa, hắn ngồi im nghe tim mình đập chậm lại từng nhịp một. ", 4)
	m := map[string]any{
		"chapter":   1,
		"beat":      1,
		"beat_goal": "Mở chương: thiết lập khủng hoảng nợ máu",
		"variants": []map[string]any{
			{"focus": "noi_tam", "content": longVi, "score": 72, "note": "sâu nhưng chậm"},
			{"focus": "doi_thoai", "content": longVi, "score": 81, "note": "căng, có ngầm ý"},
			{"focus": "hanh_dong", "content": longVi, "score": 77, "note": "nhanh, hơi mỏng"},
		},
		"final_content":   longVi + "Câu chốt buông xuống như một nhát dao lặng lẽ.",
		"merge_rationale": "Lấy khung đối thoại bản B, chèn hai nhịp nội tâm bản A, giữ cú chốt hành động bản C",
	}
	if mutate != nil {
		mutate(m)
	}
	raw, err := json.Marshal(m)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	return raw
}

func newBeatStore(t *testing.T) *store.Store {
	t.Helper()
	s := store.NewStore(t.TempDir())
	if err := s.Init(); err != nil {
		t.Fatalf("Init: %v", err)
	}
	if err := s.Progress.Init("test", 10); err != nil {
		t.Fatalf("Progress.Init: %v", err)
	}
	return s
}

func TestDraftBeatHappyPathAppendsAndRecords(t *testing.T) {
	s := newBeatStore(t)
	tool := NewDraftBeatTool(s)

	out, err := tool.Execute(context.Background(), beatArgs(t, nil))
	if err != nil {
		t.Fatalf("Execute: %v", err)
	}
	var res map[string]any
	if err := json.Unmarshal(out, &res); err != nil {
		t.Fatalf("Unmarshal result: %v", err)
	}
	if res["written"] != true || res["gate_passed"] != true {
		t.Fatalf("expected written+gate_passed, got %v", res)
	}
	if res["best_variant"] != "doi_thoai" {
		t.Fatalf("best_variant should be doi_thoai (score 81), got %v", res["best_variant"])
	}

	draft, err := s.Drafts.LoadDraft(1)
	if err != nil || draft == "" {
		t.Fatalf("draft should exist after beat append, err=%v", err)
	}
	if !strings.Contains(draft, "nhát dao lặng lẽ") {
		t.Fatalf("draft should contain final_content")
	}

	// Hồ sơ beat phải tồn tại và gate_passed=true
	var rec map[string]any
	data, err := s.Drafts.LoadBeatRecordRaw(1, 1)
	if err != nil {
		t.Fatalf("beat record missing: %v", err)
	}
	if err := json.Unmarshal(data, &rec); err != nil {
		t.Fatalf("beat record not json: %v", err)
	}
	if rec["gate_passed"] != true {
		t.Fatalf("beat record gate_passed should be true")
	}

	// Beat thứ hai append tiếp, draft dài ra
	out2, err := tool.Execute(context.Background(), beatArgs(t, func(m map[string]any) {
		m["beat"] = 2
		m["final_content"] = strings.Repeat("Trận mưa đêm dội xuống mái tôn, từng hạt gõ vào giấc ngủ chưa kịp đến của hắn. ", 3)
	}))
	if err != nil {
		t.Fatalf("Execute beat 2: %v", err)
	}
	var res2 map[string]any
	_ = json.Unmarshal(out2, &res2)
	full, _ := s.Drafts.LoadDraft(1)
	if !strings.Contains(full, "mái tôn") || !strings.Contains(full, "nhát dao") {
		t.Fatalf("draft should contain both beats")
	}
}

func TestDraftBeatGateRejectsEmDashAndEnglish(t *testing.T) {
	s := newBeatStore(t)
	tool := NewDraftBeatTool(s)

	// Em dash trong bản cuối → gate chặn, draft KHÔNG được ghi
	_, err := tool.Execute(context.Background(), beatArgs(t, func(m map[string]any) {
		m["final_content"] = strings.Repeat("Hắn bước tới — chậm rãi và lạnh lùng, như thể thời gian không còn nghĩa lý gì nữa. ", 3)
	}))
	if err == nil || !strings.Contains(err.Error(), "KHÔNG QUA GATE") {
		t.Fatalf("expected gate rejection for em dash, got %v", err)
	}
	if draft, _ := s.Drafts.LoadDraft(1); draft != "" {
		t.Fatalf("draft must stay empty after gate fail, got %q", draft)
	}
	// Hồ sơ FAIL vẫn được lưu để truy vết
	if data, err := s.Drafts.LoadBeatRecordRaw(1, 1); err == nil {
		var rec map[string]any
		_ = json.Unmarshal(data, &rec)
		if rec["gate_passed"] != false {
			t.Fatalf("fail record should have gate_passed=false")
		}
	} else {
		t.Fatalf("fail record should still be saved: %v", err)
	}

	// Tiếng Anh lọt → gate chặn
	_, err = tool.Execute(context.Background(), beatArgs(t, func(m map[string]any) {
		m["final_content"] = strings.Repeat("Hắn nhìn suddenly ra ngoài cửa sổ, lòng nặng trĩu những điều chưa kịp nói thành lời. ", 3)
	}))
	if err == nil || !strings.Contains(err.Error(), "KHÔNG QUA GATE") {
		t.Fatalf("expected gate rejection for english word, got %v", err)
	}
}

func TestDraftBeatValidatesVariants(t *testing.T) {
	s := newBeatStore(t)
	tool := NewDraftBeatTool(s)

	// Thiếu bản (2/3)
	_, err := tool.Execute(context.Background(), beatArgs(t, func(m map[string]any) {
		v := m["variants"].([]map[string]any)
		m["variants"] = v[:2]
	}))
	if err == nil || !strings.Contains(err.Error(), "đúng 3 bản") {
		t.Fatalf("expected 3-variant requirement, got %v", err)
	}

	// Focus trùng
	_, err = tool.Execute(context.Background(), beatArgs(t, func(m map[string]any) {
		v := m["variants"].([]map[string]any)
		v[1]["focus"] = "noi_tam"
	}))
	if err == nil || !strings.Contains(err.Error(), "trùng lặp") {
		t.Fatalf("expected duplicate focus rejection, got %v", err)
	}

	// Bản quá ngắn (làm phép)
	_, err = tool.Execute(context.Background(), beatArgs(t, func(m map[string]any) {
		v := m["variants"].([]map[string]any)
		v[0]["content"] = "Ngắn quá."
	}))
	if err == nil || !strings.Contains(err.Error(), "quá ngắn") {
		t.Fatalf("expected short-variant rejection, got %v", err)
	}

	// Score ngoài khoảng
	_, err = tool.Execute(context.Background(), beatArgs(t, func(m map[string]any) {
		v := m["variants"].([]map[string]any)
		v[2]["score"] = 120
	}))
	if err == nil || !strings.Contains(err.Error(), "0-100") {
		t.Fatalf("expected score range rejection, got %v", err)
	}

	// Thiếu merge_rationale
	_, err = tool.Execute(context.Background(), beatArgs(t, func(m map[string]any) {
		m["merge_rationale"] = "  "
	}))
	if err == nil || !strings.Contains(err.Error(), "merge_rationale") {
		t.Fatalf("expected merge_rationale requirement, got %v", err)
	}
}

func TestDraftBeatRespectsCompletedChapter(t *testing.T) {
	s := newBeatStore(t)
	if err := s.Progress.MarkChapterComplete(1, 3000, "", ""); err != nil {
		t.Fatalf("MarkChapterComplete: %v", err)
	}
	tool := NewDraftBeatTool(s)
	out, err := tool.Execute(context.Background(), beatArgs(t, nil))
	if err != nil {
		t.Fatalf("Execute: %v", err)
	}
	var res map[string]any
	_ = json.Unmarshal(out, &res)
	if res["skipped"] != true {
		t.Fatalf("completed chapter (not in rewrite queue) should be skipped, got %v", res)
	}
}
