# Changelog

Tất cả thay đổi đáng chú ý của dự án được ghi lại ở đây.

Định dạng dựa trên [Keep a Changelog](https://keepachangelog.com/),
và dự án tuân theo [Semantic Versioning](https://semver.org/).

## [1.0.0] - 2026-06-29

### Added
- Lần phát hành chính thức đầu tiên của **OUMStudio-Novel**.
- Engine multi-agent autonomous (Coordinator / Architect / Writer / Editor) — fork tiếng Việt của `voocel/ainovel-cli`.
- Store-first + checkpoint resume cho tiểu thuyết rất dài.
- Context compaction (ctxpack) cho truyện hàng trăm chương.
- Lớp chất lượng prose tiếng Việt: cổng chặn cứng 0 em dash / 0 tiếng Anh.
- 30+ rule liền mạch (continuity) tiếng Việt.
- 9 style Việt: ngôn tình, tiên hiệp, đô thị, trinh thám, cung đấu, fantasy, romance, suspense, default.
- Hội đồng phê bình Nacharium (5 góc nhanh / 13 góc đầy đủ).
- Xuất EPUB premium dark/bright adaptive + TXT.
- Import truyện cũ → phân tích → viết tiếp.
- Diag / observability + theo dõi usage và chi phí token.
- `oumstudio/oum-prose-verify.py` — script kiểm tra prose (`--json`, `--strict`).
- Tài liệu: README, architecture, context-management, observability.
- License Apache 2.0 + NOTICE attribution.

### Testing
- 19 package pass, 0 FAIL (`go test ./...`).

[1.0.0]: https://github.com/NachaFromMars/oumstudio-novel/releases/tag/v1.0.0
