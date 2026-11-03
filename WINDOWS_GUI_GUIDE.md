# Hướng dẫn sử dụng ipadecrypt-webGUI (Windows)

**ipadecrypt-webGUI** là giao diện đồ họa nền web hiện đại cho [ipadecrypt](https://github.com/londek/ipadecrypt), hỗ trợ giải mã các ứng dụng iOS/iPadOS tải trực tiếp từ App Store hoặc file IPA có sẵn, quản lý nhiều tài khoản Apple ID, và kết nối với thiết bị Jailbreak để tạo chứng chỉ FairPlay 1 lần duy nhất.

---

## 🚀 Khởi chạy nhanh

1. Nhấp đúp chuột vào file **`ipadecrypt-gui.bat`** (hoặc chạy lệnh `ipadecrypt.exe gui`).
2. Trình duyệt mặc định sẽ tự động mở địa chỉ `http://127.0.0.1:54322`.
3. Giao diện webGUI sẵn sàng sử dụng!

---

## 🌟 Các tính năng chính

### 1. Quản lý nhiều tài khoản Apple ID (Multi-Account)
- Hỗ trợ đăng nhập và lưu trữ **nhiều tài khoản Apple ID**.
- Hỗ trợ xác thực 2 bước (**2FA**) gửi mã về thiết bị Apple.
- Chuyển đổi tài khoản đang hoạt động chỉ với 1 cú nhấp chuột.
- Trạng thái trực quan: hiển thị rõ tài khoản nào đã được tạo chứng chỉ FairPlay (màu xanh) hoặc chưa tạo (màu vàng).
- Tự động đồng bộ với cấu hình CLI của ipadecrypt.

### 2. Tìm kiếm ứng dụng trên App Store
- Tìm kiếm theo tên ứng dụng, từ khóa, hoặc Bundle ID (ví dụ `com.spotify.client`).
- Lựa chọn quốc gia App Store (Việt Nam, Mỹ, Nhật Bản, v.v.).
- Xem thông tin ứng dụng: Icon, tên, nhà phát triển, phiên bản, giá, nền tảng (iOS, iPadOS).
- Tùy chọn **"Tải & Giải mã"** trực tiếp với 1 cú nhấp chuột.
- Tùy chọn **"Xem phiên bản"** để xem toàn bộ lịch sử các bản phát hành cũ.

### 3. Giải mã trực tiếp & Giải mã file IPA cục bộ
- **Từ App Store**: Nhập Bundle ID, App ID hoặc URL App Store để tự động tải về và giải mã.
- **Từ máy tính**: Kéo thả hoặc chọn file `.ipa` đã mã hóa trên máy tính để giải mã ngay lập tức.
- Các tùy chọn nâng cao:
  - Bắt buộc giải mã trên thiết bị (`--on-device`)
  - Không sử dụng thiết bị (`--no-on-device`)
  - Mở rộng hỗ trợ dòng máy (`--patch-device-type`)
  - Bỏ qua các tiện ích mở rộng (`--skip-appex`)

### 4. Lịch sử phiên bản App Store
- Tra cứu danh sách toàn bộ các phiên bản đã phát hành trên App Store của ứng dụng.
- Hiển thị Version ID, phiên bản marketing, ngày phát hành.
- Cho phép giải mã bản cũ tùy ý.

### 5. Thiết bị Jailbreak & Tạo chứng chỉ FairPlay (Mint Credential)
- Kết nối qua mạng **Wi-Fi** (IP + Cổng SSH 22) hoặc qua **cáp USB** (thông qua `iproxy`).
- Nút **"Kiểm tra kết nối"** tự động kiểm tra:
  - Kết nối OpenSSH với quyền `root`
  - Gói `AppSync Unified`
  - Công cụ `appinst`
  - Helper giải mã trên máy
- Nút **"Tạo chứng chỉ FairPlay (Mint)"**: Kết nối trích xuất khóa giải mã 16-byte duy nhất 1 lần cho tài khoản Apple ID. Sau bước này, việc giải mã các ứng dụng arm64 sẽ diễn ra hoàn toàn offline trên máy tính mà không cần cắm iPhone/iPad nữa!

### 6. Quản lý tệp IPA đã giải mã
- Danh sách các file IPA đã giải mã lưu tại thư mục `./decrypted`.
- Nút **"Mở thư mục trong Explorer"** để truy cập ngay thư mục chứa file.
- Nút **"Tải về máy"** để tải file IPA qua trình duyệt.
- Xóa file trực tiếp từ giao diện.

### 7. Cài đặt vào thiết bị iOS
- Tích hợp công cụ `libimobiledevice` (`ideviceinstaller`).
- Nhận diện các thiết bị iPhone/iPad đang cắm qua USB.
- Chọn file IPA đã giải mã và nhấn **"Cài đặt vào máy"**.

---

## 🛠 Hướng dẫn biên dịch từ mã nguồn

Nếu muốn biên dịch lại file `ipadecrypt.exe`:
1. Cài đặt Go (phiên bản 1.25 trở lên).
2. Nhấp đúp vào **`build.bat`**.
3. File `ipadecrypt.exe` mới với giao diện webGUI nhúng sẵn sẽ được tạo ra tự động!
