# Contributing to OUMStudio-Novel

Cảm ơn bạn quan tâm đến OUMStudio-Novel. Tài liệu này mô tả quy ước đóng góp.

## Yêu cầu môi trường

- Go ≥ 1.25
- Python 3 (cho `oumstudio/oum-prose-verify.py`)

## Build & test

```bash
go build -o bin/oum-novel ./cmd/ainovel-cli/
go test ./...
```

Mọi PR phải giữ `go test ./...` xanh (hiện 19 package pass).

## Quy ước code

- Theo `gofmt` / `go vet` chuẩn.
- Logic engine đặt trong `internal/`; entry point trong `cmd/`.
- Assets (prompts, references, rules, styles) đặt trong `assets/` và được embed.
- Lớp prose tiếng Việt đặt trong `oumstudio/`.

## Quy ước prose tiếng Việt (bắt buộc cho nội dung Việt)

- **0 em dash** (—) trong prose.
- **0 tiếng Anh** lọt vào prose tiếng Việt.
- **0 markdown** trong văn bản truyện.
- Tôn trọng 30+ rule liền mạch trong `assets/references/continuity-rules-vi.md`.

Kiểm tra trước khi commit nội dung:

```bash
python3 oumstudio/oum-prose-verify.py --strict <file>
```

## Pull request

1. Fork + tạo branch `feature/...` hoặc `fix/...`.
2. Commit nhỏ, message rõ ràng.
3. Đảm bảo test xanh.
4. Mở PR mô tả thay đổi + lý do.

## License

Đóng góp của bạn sẽ được phát hành dưới [Apache License 2.0](LICENSE).
