# Changelog

Tất cả thay đổi đáng chú ý của dự án được ghi lại ở đây.
Định dạng dựa trên [Keep a Changelog](https://keepachangelog.com/), tuân theo [Semantic Versioning](https://semver.org/).

## [1.1.0] - 2026-06-29

### Added — Omni Novel Suite
- Hợp nhất engine autonomous + **10 skill quy trình** vào một repo duy nhất (`skills/`).
- Bundle: cw-brainstorming, cw-prose-writing, cw-story-critique, cw-official-docs, cw-style-skill-creator, cw-router, novel-guardian, novel-master, novelcore-ai, forge-novel-guard.
- `SKILLS.md` — bản đồ pipeline 7 bước (brainstorm → style → forge → scan → critique → verify → wiki → export).
- README cập nhật thành Omni Suite (bảng so sánh engine vs skill, kiến trúc hợp nhất).
- NOTICE bổ sung attribution sub-component (novel-guardian MIT).

## [1.0.0] - 2026-06-29

### Added
- Lần phát hành chính thức đầu tiên của **OUMStudio-Novel**.
- Engine multi-agent autonomous (Coordinator / Architect / Writer / Editor) — fork tiếng Việt của `voocel/ainovel-cli`.
- Store-first + checkpoint resume cho tiểu thuyết rất dài.
- Context compaction (ctxpack) cho truyện hàng trăm chương.
- Lớp chất lượng prose tiếng Việt: cổng chặn cứng 0 em dash / 0 tiếng Anh.
- 30+ rule liền mạch (continuity) tiếng Việt.
- 9 style Việt + hội đồng phê bình Nacharium.
- Xuất EPUB premium dark/bright adaptive + TXT.
- Import truyện cũ → phân tích → viết tiếp.
- License Apache 2.0 + NOTICE attribution.

### Testing
- 19 package pass, 0 FAIL (`go test ./...`).

[1.1.0]: https://github.com/NachaFromMars/oumstudio-novel/releases/tag/v1.1.0
[1.0.0]: https://github.com/NachaFromMars/oumstudio-novel/releases/tag/v1.0.0
