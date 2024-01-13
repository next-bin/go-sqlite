cp ccgo_windows_amd64.go ccgo_windows_arm64.go
sed -i 's/amd64/arm64/' ccgo_windows_arm64.go
cp internal/minigzip/ccgo_windows_amd64.go internal/minigzip/ccgo_windows_arm64.go
sed -i 's/amd64/arm64/' internal/minigzip/ccgo_windows_arm64.go
cp internal/example/ccgo_windows_amd64.go internal/example/ccgo_windows_arm64.go
sed -i 's/amd64/arm64/' internal/example/ccgo_windows_arm64.go
rm -rf include/windows/arm64/
cp -r include/windows/amd64/ include/windows/arm64/
