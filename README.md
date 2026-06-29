# OUMStudio-Novel — Omni Novel Suite

> Bộ công cụ viết tiểu thuyết tiếng Việt **trọn vẹn**: engine autonomous + lớp chất lượng prose Việt + bộ skill quy trình. Một câu yêu cầu → tiểu thuyết hoàn chỉnh, chất lượng văn Việt premium.

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](LICENSE)
[![Go](https://img.shields.io/badge/Go-1.25+-00ADD8?logo=go)](https://go.dev)
[![Tests](https://img.shields.io/badge/tests-19_packages_passing-success)](#kiểm-thử)
[![Skills](https://img.shields.io/badge/novel_skills-11_bundled-orange)](#bộ-skill-novel-tích-hợp)
[![Status](https://img.shields.io/badge/status-production-brightgreen)](#)

OUMStudio-Novel hợp nhất **hai hướng** thành một bộ duy nhất:

1. **Engine autonomous** (Go, multi-agent) — *cỗ máy chạy thật*: nhập 1 câu → tự gọi tools, tự viết, tự checkpoint, tự xuất sách. Fork tiếng Việt của [`voocel/ainovel-cli`](https://github.com/voocel/ainovel-cli).
2. **Bộ skill chất lượng** — *quy trình kiểm soát*: continuity 30+ rule, văn phong Việt, hội đồng phê bình, EPUB premium, brainstorm/critique/wiki.

Engine lo phần **chạy một mình**; bộ skill lo phần **chất lượng văn Việt**. Gộp lại = "Skill viết sách VN" thực thụ.

---

## Mục lục

1. [Tại sao là Omni Suite](#tại-sao-là-omni-suite)
2. [Kiến trúc](#kiến-trúc)
3. [Bộ skill novel tích hợp](#bộ-skill-novel-tích-hợp)
4. [Cài đặt](#cài-đặt)
5. [Sử dụng nhanh](#sử-dụng-nhanh)
6. [Lớp chất lượng prose Việt](#lớp-chất-lượng-prose-việt)
7. [Style có sẵn](#style-có-sẵn)
8. [Cấu hình](#cấu-hình)
9. [Kiểm thử](#kiểm-thử)
10. [Cấu trúc thư mục](#cấu-trúc-thư-mục)
11. [Attribution & License](#attribution--license)

---

## Tại sao là Omni Suite

| Khía cạnh | Engine autonomous | Bộ skill chất lượng |
|-----------|-------------------|---------------------|
| Bản chất | Cỗ máy chạy thật (binary) | Quy trình + ruleset + prompt |
| Khởi động → để chạy → lấy output | ✅ mạnh | — |
| Checkpoint resume cấp tool-call | ✅ | — |
| Context compaction 4 tầng | ✅ | — |
| Can thiệp realtime | ✅ | — |
| Continuity 30+ rule cứng | ✅ (đã nhúng) + 🔧 novel-guardian/novel-master | ✅ |
| Văn phong Việt + chống AI | ✅ (gate) + 🔧 cw-prose-writing | ✅ |
| Hội đồng phê bình đa góc | ✅ Nacharium | 🔧 cw-story-critique |
| Brainstorm / wiki / style guide | — | 🔧 cw-* |
| EPUB premium | ✅ | 🔧 forge-novel-guard |

→ Không thay thế nhau. **Gộp = cả hai.**

## Kiến trúc

```
┌─────────────┐   điều phối   ┌─────────────┐
│ Coordinator │──────────────▶│  Architect  │  premise / outline / cast / world
└─────────────┘               └─────────────┘
       │                             │
       ▼                             ▼
┌─────────────┐  draft→store  ┌─────────────┐  review/commit/continuity
│   Writer    │──────────────▶│   Editor    │
└─────────────┘               └─────────────┘
       │                             │
       └──────────┬──────────────────┘
                  ▼
   ┌──────────────────────────────────┐
   │  Lớp chất lượng (oumstudio/ + skills/)  │
   │  prose-verify · Nacharium · 30 rule     │
   │  cw-* · novel-guardian · novel-master   │
   └──────────────────────────────────┘
                  ▼
        Store (checkpoint) → Export (EPUB/TXT)
```

Chi tiết: [`docs/architecture.md`](docs/architecture.md), [`docs/context-management.md`](docs/context-management.md), [`docs/observability.md`](docs/observability.md).

## Bộ skill novel tích hợp

11 skill lõi (engine + lớp prose + 10 skill quy trình) nằm trong [`skills/`](skills/) và [`oumstudio/`](oumstudio/). Xem [`SKILLS.md`](SKILLS.md) cho bản đồ đầy đủ.

| Skill | Vai trò trong pipeline |
|-------|------------------------|
| **oumstudio** (engine layer) | Cổng prose 0 em dash/0 tiếng Anh + hội đồng Nacharium |
| **cw-brainstorming** | PRE-FORGE: outline, beats, character arc |
| **cw-style-skill-creator** | SETUP: tạo style guide cho project |
| **cw-prose-writing** | FORGE: viết prose theo style guide |
| **cw-story-critique** | POST-FORGE: phê bình pacing/nhân vật/thoại |
| **cw-official-docs** | WIKI: profile nhân vật, lore, địa danh |
| **cw-router** | DISPATCH: tự chọn skill cho đúng task |
| **novel-guardian** | SCAN: 10 rule liền mạch (T/C/W/P/V/D), pacing, style |
| **novel-master** | CHECK: 5 lớp (bible/continuity/prose/pace/voice), 30+ rule |
| **novelcore-ai** | FORGE-OPT: prompt 4 khối, 3 bản, token budget, context phân vùng |
| **forge-novel-guard** | VERIFY: prose/epub/send guard (4 script) |

## Cài đặt

### Yêu cầu
- **Go ≥ 1.25** (build engine)
- **Node.js** (cho novel-guardian / novel-master CLI)
- **Python 3** (cho verify scripts)

### Build engine
```bash
git clone https://github.com/NachaFromMars/oumstudio-novel.git
cd oumstudio-novel
go build -o bin/oum-novel ./cmd/ainovel-cli/
```

### Docker
```bash
docker compose up --build
```

## Sử dụng nhanh

Chạy TUI:
```bash
./bin/oum-novel
```
Hoặc headless (non-interactive):
```bash
./bin/oum-novel --headless --prompt "Viết truyện tiên hiệp 30 chương về..."
```

## Lớp chất lượng prose Việt

```bash
python3 oumstudio/oum-prose-verify.py --strict chuong-01.txt    # phát hiện em dash, tiếng Anh, markdown, AI trace
node skills/novel-guardian/scripts/*.mjs                         # scan continuity
node skills/novel-master/scripts/novel-master.mjs ...            # check 5 lớp
python3 skills/forge-novel-guard/scripts/master-verify.py ...    # verify trước khi giao
```

## Style có sẵn

| Style | File |
|-------|------|
| Ngôn tình | `assets/styles/ngon-tinh.md` |
| Tiên hiệp | `assets/styles/tien-hiep.md` |
| Đô thị | `assets/styles/do-thi.md` |
| Trinh thám | `assets/styles/trinh-tham.md` |
| Cung đấu | `assets/styles/cung-dau.md` |
| Fantasy / Romance / Suspense / Default | `assets/styles/*.md` |

## Cấu hình

Provider + roles trong `~/.ainovel/config.json`. Mẫu: [`internal/bootstrap/config.example.jsonc`](internal/bootstrap/config.example.jsonc).

> ⚠️ Tránh model bỏ tham số `temperature`.

## Kiểm thử

```bash
go test ./...
```
**19 package pass, 0 FAIL.**

## Cấu trúc thư mục

```
.
├── cmd/ainovel-cli/      # entry point
├── internal/             # engine core
├── assets/               # prompts/references/rules/styles (embed)
├── oumstudio/            # lớp prose Việt (verify + Nacharium council + SKILL)
├── skills/               # 10 skill quy trình (cw-*, novel-guardian, novel-master, ...)
├── docs/                 # architecture, context, observability
├── SKILLS.md             # bản đồ pipeline skill
├── LICENSE / NOTICE      # Apache 2.0 + attribution
└── README.md
```

## Attribution & License

Phát hành dưới **[Apache License 2.0](LICENSE)**.

- Engine là fork phái sinh của [`voocel/ainovel-cli`](https://github.com/voocel/ainovel-cli) — attribution trong [`NOTICE`](NOTICE).
- `skills/novel-guardian/` giữ LICENSE riêng (MIT, tương thích) của nó.
- Lớp prose Việt, Nacharium, các style Việt, bộ skill quy trình © Tiểu Tâm × Nấng.
