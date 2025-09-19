.PHONY: add sub

# add関数のアセンブリを表示（最適化なし）
add:
	@echo "=== add関数のアセンブリ（最適化なし -N -l）==="
	@GOARCH=amd64 go tool compile -N -l -S main.go 2>/dev/null | grep -A 15 'main.add STEXT'

# sub関数のアセンブリを表示（最適化なし）
sub:
	@echo "=== sub関数のアセンブリ（最適化なし -N -l）==="
	@GOARCH=amd64 go tool compile -N -l -S main.go 2>/dev/null | grep -A 15 'main.sub STEXT'