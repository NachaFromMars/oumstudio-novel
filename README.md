# OUMStudio-Novel

> Engine viết tiểu thuyết dài **autonomous** cho tiếng Việt — multi-agent, store-first, checkpoint-resume, EPUB premium.

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](LICENSE)
[![Go](https://img.shields.io/badge/Go-1.25+-00ADD8?logo=go)](https://go.dev)
[![Tests](https://img.shields.io/badge/tests-19_packages_passing-success)](#kiểm-thử)
[![Status](https://img.shields.io/badge/status-production-brightgreen)](#)

OUMStudio-Novel là một **cỗ máy chạy thật** (binary Go), không phải prompt-guide. Nhập **một câu yêu cầu** → engine tự dựng premise, outline, nhân vật, thế giới → viết từng chương → kiểm tra nhất quán → xuất **EPUB/TXT**. Hỗ trợ tiểu thuyết rất dài qua checkpoint resume, can thiệp realtime, và import truyện cũ để viết tiếp.

Engine là một **fork tiếng Việt** của [`voocel/ainovel-cli`](https://github.com/voocel/ainovel-cli), được nâng cấp với một lớp chất lượng prose tiếng Việt premium: cổng chặn cứng (0 em dash / 0 tiếng Anh), 30+ rule liền mạch, 9 style Việt, và hội đồng phê bình Nacharium.

---

## Mục lục

1. [Tính năng nổi bật](#tính-năng-nổi-bật)
2. [Kiến trúc](#kiến-trúc)
3. [Cài đặt](#cài-đặt)
4. [Sử dụng nhanh](#sử-dụng-nhanh)
5. [Lớp chất lượng prose Việt](#lớp-chất-lượng-prose-việt)
6. [Style có sẵn](#style-có-sẵn)
7. [Cấu hình](#cấu-hình)
8. [Kiểm thử](#kiểm-thử)
9. [Cấu trúc thư mục](#cấu-trúc-thư-mục)
10. [Attribution & License](#attribution--license)

---

## Tính năng nổi bật

- **Autonomous end-to-end** — 1 prompt → tiểu thuyết hoàn chỉnh, không cần can thiệp thủ công trong quá trình viết.
- **Multi-agent** — Coordinator điều phối → Architect dựng nền → Writer viết → Editor thẩm định.
- **Store-first + checkpoint** — mọi trạng thái (premise, outline, cast, world, drafts, summaries) lưu xuống store; dừng/khôi phục bất kỳ lúc nào cho truyện rất dài.
- **Context compaction** — tự nén ngữ cảnh (ctxpack) để viết hàng trăm chương mà không tràn context.
- **Cổng chất lượng prose Việt** — chặn cứng em dash, tiếng Anh lọt vào prose tiếng Việt.
- **30+ rule liền mạch** — kiểm tra continuity (timeline, teleportation, power-jump, vật phẩm, đại từ...).
- **9 style Việt** — ngôn tình, tiên hiệp, đô thị, trinh thám, cung đấu, fantasy, romance, suspense, default.
- **Hội đồng Nacharium** — review đa góc nhìn (5 góc nhanh / 13 góc đầy đủ) cho chương then chốt.
- **EPUB premium** — xuất EPUB dark/bright adaptive + TXT.
- **Import truyện cũ** — phân tích chương có sẵn → dựng foundation → viết tiếp.
- **Quan sát & chẩn đoán** — diag/observability tích hợp, theo dõi usage + chi phí token.

## Kiến trúc

```
┌─────────────┐   điều phối   ┌─────────────┐
│ Coordinator │──────────────▶│  Architect  │  dựng premise / outline / cast / world
└─────────────┘               └─────────────┘
       │                             │
       │                             ▼
       │                       ┌─────────────┐
       │       viết chương     │   Writer    │  draft → store
       ├──────────────────────▶└─────────────┘
       │                             │
       │                             ▼
       │       thẩm định        ┌─────────────┐
       └──────────────────────▶│   Editor    │  review / commit / continuity
                                └─────────────┘
                                       │
                          ┌────────────┴────────────┐
                          ▼                         ▼
                   Store (checkpoint)        Export (EPUB/TXT)
```

Tham khảo chi tiết: [`docs/architecture.md`](docs/architecture.md), [`docs/context-management.md`](docs/context-management.md), [`docs/observability.md`](docs/observability.md).

## Cài đặt

### Yêu cầu

- **Go ≥ 1.25** để build từ source.
- Một provider LLM (cấu hình trong `~/.ainovel/config.json`).

### Build từ source

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

OUMStudio-Novel chạy ở chế độ **TUI** (terminal UI). Khởi động rồi nhập yêu cầu truyện trong ô nhập liệu:

```bash
./bin/oum-novel
```

Trong TUI:
- Nhập 1 câu mô tả truyện → engine tự dựng nền và bắt đầu viết.
- Theo dõi tiến độ realtime, can thiệp qua command palette.
- Xuất EPUB/TXT khi đã có chương hoàn thành.

> Lưu ý: phiên bản này **không** nhận yêu cầu truyện trực tiếp qua tham số dòng lệnh; phải nhập trong TUI.

## Lớp chất lượng prose Việt

Đây là phần OUMStudio bổ sung lên trên engine gốc:

- **`oumstudio/oum-prose-verify.py`** — script kiểm tra prose: phát hiện em dash, tiếng Anh lọt vào, markdown, dấu vết AI; hỗ trợ `--json` và `--strict`.

  ```bash
  python3 oumstudio/oum-prose-verify.py --strict chuong-01.txt
  python3 oumstudio/oum-prose-verify.py --json chuong-01.txt
  ```

- **`oumstudio/nacharium-council.md`** — hội đồng phê bình Nacharium 13 góc (bộ nhanh 5 góc mặc định) để thẩm định sâu chương/bản thảo.

- **30+ rule liền mạch** — nhúng trong `assets/references/continuity-rules-vi.md`.

## Style có sẵn

| Style | File | Mô tả |
|-------|------|-------|
| Ngôn tình | `assets/styles/ngon-tinh.md` | Tình cảm, cảm xúc nội tâm |
| Tiên hiệp | `assets/styles/tien-hiep.md` | Tu tiên, công pháp, cảnh giới |
| Đô thị | `assets/styles/do-thi.md` | Hiện đại, đời thường |
| Trinh thám | `assets/styles/trinh-tham.md` | Suy luận, manh mối |
| Cung đấu | `assets/styles/cung-dau.md` | Hậu cung, mưu lược |
| Fantasy | `assets/styles/fantasy.md` | Kỳ ảo phương Tây |
| Romance | `assets/styles/romance.md` | Lãng mạn |
| Suspense | `assets/styles/suspense.md` | Hồi hộp, căng thẳng |
| Default | `assets/styles/default.md` | Mặc định |

## Cấu hình

Engine đọc cấu hình provider + roles từ `~/.ainovel/config.json`. Xem mẫu trong [`internal/bootstrap/config.example.jsonc`](internal/bootstrap/config.example.jsonc) và [`rules.md.example`](rules.md.example).

> ⚠️ Tránh model bỏ tham số `temperature` (một số proxy báo "temperature deprecated"). Model hỗ trợ `temperature` hoạt động ổn định.

## Kiểm thử

```bash
go test ./...
```

Hiện tại: **19 package pass, 0 FAIL**.

## Cấu trúc thư mục

```
.
├── cmd/ainovel-cli/      # entry point
├── internal/             # engine core (agents, host, store, tools, rules, diag, ...)
├── assets/               # prompts, references, rules, styles (embed)
├── oumstudio/            # lớp chất lượng prose Việt (verify, council, skill)
├── docs/                 # architecture, context-management, observability
├── scripts/              # tiện ích (wordcount, install, media)
├── Dockerfile
├── docker-compose.yml
├── LICENSE               # Apache 2.0
├── NOTICE                # attribution
└── README.md
```

## Attribution & License

OUMStudio-Novel được phát hành dưới **[Apache License 2.0](LICENSE)**.

Engine là một fork phái sinh của [`voocel/ainovel-cli`](https://github.com/voocel/ainovel-cli). Các thông báo bản quyền và attribution gốc được giữ trong file [`NOTICE`](NOTICE) theo yêu cầu của Apache 2.0.

Lớp chất lượng prose tiếng Việt, hội đồng Nacharium, 9 style Việt, và các nâng cấp continuity © Tiểu Tâm × Nấng.
