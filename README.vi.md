# ipadecrypt-webGUI

<p align="center">
  <img src="cmd/ipadecrypt/gui_assets/icon.png" alt="ipadecrypt logo" width="128" height="128" onerror="this.style.display='none'"/>
</p>

<p align="center">
  <strong>Bộ công cụ Giao diện Đồ họa Web (WebGUI) & Dòng lệnh (CLI) hiện đại giúp giải mã (decrypt) file IPA bảo vệ bởi FairPlay DRM từ App Store.</strong>
</p>

<p align="center">
  <a href="README.md"><strong>English</strong></a> •
  <a href="README.vi.md"><strong>Tiếng Việt</strong></a>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go" alt="Phiên bản Go"/>
  <img src="https://img.shields.io/badge/N%E1%BB%81n%20t%E1%BA%A3ng-Windows%20%7C%20macOS%20%7C%20Linux-blue?style=flat" alt="Nền tảng"/>
  <img src="https://img.shields.io/badge/Gi%E1%BA%A5y%20ph%C3%A9p-MIT-green?style=flat" alt="Giấy phép"/>
  <img src="https://img.shields.io/badge/FairPlay-arm64%20Offline%20Decrypt-success?style=flat" alt="FairPlay Giải mã ngoại tuyến"/>
  <img src="https://img.shields.io/badge/Giao%20di%E1%BB%87n-Web%20Interface%20%26%20CLI-orange?style=flat" alt="Giao diện Web"/>
</p>

---

## 📖 Giới thiệu tổng quan

**ipadecrypt-webGUI** là phiên bản giao diện đồ họa nền web và gói tiện ích mở rộng toàn diện được xây dựng dựa trên [ipadecrypt](https://github.com/londek/ipadecrypt). Bộ công cụ này hỗ trợ các nhà nghiên cứu bảo mật, kỹ sư dịch ngược (reverse engineers) và lập trình viên iOS dễ dàng tải các ứng dụng nguyên bản từ App Store và giải mã triệt để lớp bảo vệ bản quyền FairPlay DRM.

### Điểm vượt trội so với các phương pháp dump app truyền thống

Trước đây, khi cần trích xuất file IPA đã giải mã (thông qua các công cụ như *frida-ios-dump*, *Clutch*, *bfdecrypt*), quy trình bắt buộc phải:
1. Cài đặt ứng dụng vào một iPhone/iPad đã Jailbreak.
2. Mở ứng dụng để nạp các trang mã nhị phân Mach-O vào bộ nhớ RAM.
3. Chạy script hook để kết xuất (dump) các trang bộ nhớ đã giải mã ra file.

**ipadecrypt tạo nên một bước đột phá hoàn toàn mới:**
- **Tạo chứng chỉ FairPlay 1 lần duy nhất (Mint Credential)**: Chỉ cần kết nối thiết bị Jailbreak arm64 đúng **1 lần** để trích xuất khóa giải mã phần cứng 16-byte liên kết với tài khoản Apple ID của bạn.
- **Giải mã ngoại tuyến 100% (Offline arm64)**: Sau khi đã tạo chứng chỉ, mọi ứng dụng iOS 64-bit hiện đại tải về từ App Store sẽ được giải mã **ngay trên máy tính hoàn toàn offline mà không cần cắm iPhone/iPad nữa**!
- **Tự động Fallback trên thiết bị cho app 32-bit**: Đối với các ứng dụng cũ vẫn còn phân vùng 32-bit (`armv7`), hệ thống tự động đẩy helper lên máy Jailbreak để xử lý giải mã tại chỗ.
- **Giao diện Web trực quan, đầy đủ tính năng**: Tích hợp sẵn quản lý nhiều tài khoản Apple ID (hỗ trợ 2FA), tra cứu App Store đa quốc gia, xem lại lịch sử các phiên bản cũ để tải bản cũ tùy ý, và cài đặt trực tiếp file IPA vào thiết bị iOS qua bộ công cụ `libimobiledevice` tích hợp sẵn.

---

## ✨ Các tính năng nổi bật

- **🌐 Giao diện Web hiện đại & Song ngữ VI / EN**:
  - Tùy chọn giao diện Tối (Dark mode) và Sáng (Light mode) bảo vệ mắt.
  - Hỗ trợ chuyển đổi nhanh giữa Tiếng Việt và Tiếng Anh chỉ với 1 cú click.
- **👥 Quản lý nhiều tài khoản Apple ID (Multi-Account)**:
  - Lưu trữ và quản lý đồng thời nhiều tài khoản Apple ID (Việt Nam, Mỹ, Nhật, v.v.).
  - Hỗ trợ đầy đủ xác thực 2 bước (**2FA**) với hộp thoại nhập mã xác nhận 6 số từ thiết bị Apple.
  - Chuyển đổi tài khoản đang hoạt động chỉ với 1 cú nhấp chuột.
  - Huy hiệu trực quan: hiển thị rõ tài khoản đã được cấp chứng chỉ FairPlay (màu xanh lá) hay chưa (màu vàng).
- **🔍 Công cụ tìm kiếm App Store thông minh**:
  - Tra cứu theo tên ứng dụng, từ khóa, tên nhà phát triển hoặc Bundle ID (ví dụ: `com.spotify.client`).
  - Hỗ trợ chọn quốc gia Storefront (Việt Nam, Mỹ, Nhật Bản, Hàn Quốc, Anh, Đức, Pháp, Trung Quốc,...).
  - Xem nhanh thông tin ứng dụng: biểu tượng, danh mục, giá tiền, phiên bản, kiến trúc tương thích.
  - Tùy chọn nhanh: **"Tải & Giải mã"** hoặc **"Xem phiên bản"** ngay trên từng kết quả tìm kiếm.
- **🔓 Bảng điều khiển giải mã FairPlay mạnh mẽ**:
  - **Từ App Store**: Nhập Bundle ID, App ID hoặc đường link App Store URL để tự động tải về và giải mã.
  - **Từ tệp máy tính**: Kéo thả hoặc chọn trực tiếp file `.ipa` đã mã hóa lưu trên máy tính để giải mã.
  - Đầy đủ các cờ tùy chọn nâng cao theo chuẩn ipadecrypt:
    - `--external-version-id`: Chỉ định ID phiên bản cũ cần giải mã.
    - `--storefront`: Ghi đè mã quốc gia / ID storefront App Store.
    - `--skip-appex`: Bỏ qua các tiện ích mở rộng (`PlugIns/*.appex`) để tăng tốc độ xử lý.
    - `--patch-device-type`: Mở rộng `UIDeviceFamily` trong file `Info.plist` giúp app iPad có thể cài đặt được trên iPhone (hoặc ngược lại).
    - `--extra-verify`: So sánh từng byte với file gốc Mach-O để đảm bảo tính toàn vẹn tuyệt đối.
    - `--no-verify`: Bỏ qua bước kiểm tra Mach-O cryptid.
    - Chế độ giải mã linh hoạt: Tự động, Bắt buộc trên thiết bị (`--on-device`), hoặc Chỉ giải mã offline (`--no-on-device`).
  - Thanh tiến trình trực quan hiển thị tốc độ tải, dung lượng và nhật ký theo thời gian thực.
- **📜 Tra cứu lịch sử phiên bản App Store & Rollback**:
  - Truy vấn toàn bộ danh sách các phiên bản đã từng phát hành trên máy chủ Apple.
  - Hiển thị Version ID, số phiên bản phát hành, ngày cập nhật.
  - 1 cú nhấp chuột để tải và giải mã phiên bản cũ mong muốn.
- **📱 Kết nối thiết bị Jailbreak & Tạo chứng chỉ FairPlay (Mint)**:
  - Kết nối qua cáp **USB** (thông qua đường hầm `iproxy` tự động) hoặc mạng **Wi-Fi SSH** (`host:port`, `root:alpine`).
  - Nút **"Kiểm tra kết nối"** tự động rà soát 4 thành phần thiết yếu: Root OpenSSH, `AppSync Unified`, `appinst`, và Binary Helper.
  - Nút **"Tạo chứng chỉ FairPlay"**: Tự động đưa payload lên máy để trích xuất khóa giải mã 16-byte duy nhất cho tài khoản Apple ID.
- **📂 Quản lý tệp IPA đã giải mã & Cài đặt vào iPhone/iPad**:
  - Quản lý danh sách các file IPA giải mã lưu tại thư mục `./decrypted`.
  - Mở nhanh thư mục trên Windows Explorer, tải file qua trình duyệt, hoặc xóa file thừa.
  - **Cài đặt trực tiếp vào thiết bị**: Tự động phát hiện các máy iPhone/iPad đang cắm USB thông qua `idevice_id` và cài đặt file IPA trực tiếp bằng `ideviceinstaller`.

---

## 🏗️ Sơ đồ kiến trúc & Cơ chế hoạt động

```mermaid
flowchart TD
    subgraph Setup ["Thiết lập 1 lần duy nhất"]
        A[Thiết bị Jailbreak<br/>arm64 iOS 14-16] -->|Cáp USB hoặc Wi-Fi SSH| B[Tạo chứng chỉ FairPlay (Mint)]
        B -->|Trích xuất khóa 16-byte| C[(Bộ nhớ chứng chỉ cục bộ<br/>~/.ipadecrypt/)]
    end

    subgraph Operation ["Quy trình giải mã hàng ngày"]
        D[Tìm kiếm App Store / Nhập Bundle ID] --> E[Tải về từ máy chủ Apple CDN<br/>Xác thực qua AppleMediaServices]
        F[File .ipa cục bộ có sẵn] --> G{Kiểm tra kiến trúc}
        E --> G
        
        C -.->|Khóa giải mã| H[Bộ giải mã ngoại tuyến<br/>Phân vùng arm64]
        G -->|Phân vùng arm64| H
        G -->|Phân vùng 32-bit armv7| I[Bộ giải mã dự phòng trên máy<br/>ipadecrypt-helper]
        
        H --> J[Tệp IPA đã giải mã<br/>./decrypted/]
        I --> J
        
        J --> K[Tải về máy tính qua Trình duyệt]
        J --> L[Cài đặt trực tiếp vào iPhone qua USB]
    end
```

---

## 📋 Yêu cầu hệ thống & Điều kiện tiên quyết

### Máy tính sử dụng (Host PC)
- **Hệ điều hành**: Windows 10/11 (64-bit), macOS 11+, hoặc Linux.
- **Go**: Phiên bản 1.25 trở lên (chỉ cần khi muốn tự biên dịch từ mã nguồn).
- **Kết nối mạng**: Internet ổn định để tải IPA từ máy chủ Apple.
- **Driver USB Apple (riêng cho Windows)**: Cài đặt iTunes hoặc ứng dụng *Apple Devices* từ Microsoft Store để Windows nhận diện thiết bị iOS qua cổng USB.

### Thiết bị iOS Jailbreak (Chỉ cần 1 lần để Mint chứng chỉ)
- Thiết bị iOS vật lý chạy chip **arm64** đã Jailbreak (từ iPhone 6s đến iPhone X, hoặc iPad tương đương chạy iOS 14.x đến 16.x).
  - Hỗ trợ các công cụ Jailbreak phổ biến: **palera1n**, **Dopamine**, **checkra1n**, **unc0ver**, **Taurine**, v.v.
- Các gói (package) bắt buộc cài đặt trên thiết bị (qua Sileo, Zebra hoặc Cydia):
  1. **OpenSSH** (Kích hoạt SSH tại cổng 22, tài khoản `root`, mật khẩu mặc định là `alpine`).
  2. **AppSync Unified** (Nguồn repo Sileo: `https://lukezgd.github.io/repo` hoặc repo Karen).
  3. **appinst** (Công cụ cài đặt IPA qua dòng lệnh từ repo Karen/BigBoss).

> [!NOTE]
> Sau khi đã tạo chứng chỉ thành công cho tài khoản Apple ID, thiết bị Jailbreak **không còn cần thiết** cho việc giải mã các ứng dụng arm64 thông thường nữa!

---

## 🚀 Khởi chạy nhanh trên Windows

1. **Tải về hoặc Clone kho mã nguồn**:
   ```sh
   git clone https://github.com/blackcatxhan/ipadecrypt-webGUI.git
   cd ipadecrypt-webGUI
   ```

2. **Khởi động Giao diện WebGUI**:
   - Nhấp đúp chuột vào file **`ipadecrypt-gui.bat`** (hoặc chạy file PowerShell `.\ipadecrypt-gui.ps1`).
   - Hoặc mở cửa sổ dòng lệnh và gõ:
     ```cmd
     .\ipadecrypt.exe gui
     ```

3. **Mở trình duyệt**:
   - Trình duyệt mặc định trên máy sẽ tự động truy cập địa chỉ:
     ```
     http://127.0.0.1:54322
     ```
   - Các file IPA sau khi giải mã sẽ tự động lưu tại thư mục `./decrypted/`.

---

## 🛠️ Hướng dẫn biên dịch từ mã nguồn

Nếu bạn muốn tự tay biên dịch file thực thi `ipadecrypt.exe` từ mã nguồn:

### Điều kiện tiên quyết
- Cài đặt **Go 1.25+** từ trang chủ [go.dev/dl](https://go.dev/dl/).
- Đảm bảo lệnh `go` đã được nhận diện trong biến môi trường `PATH`.

### Biên dịch trên Windows
Chỉ cần nhấp đúp chuột vào file script tự động:
```cmd
build.bat
```

Hoặc chạy lệnh thủ công:
```cmd
go build -trimpath -ldflags="-s -w" -o ipadecrypt.exe ./cmd/ipadecrypt
```

### Biên dịch trên macOS / Linux
```sh
go build -trimpath -ldflags="-s -w" -o ipadecrypt ./cmd/ipadecrypt
```

*Ghi chú: Toàn bộ giao diện tĩnh (`gui_assets/`) và các tệp nhị phân mint thiết bị (`internal/credmint/bundle/`) đều được nhúng trực tiếp vào trong file thực thi duy nhất bằng tính năng `go:embed`.*

---

## 📚 Hướng dẫn sử dụng chi tiết từng bước

### Bước 1: Đăng nhập tài khoản Apple ID
1. Trên giao diện webGUI, chuyển sang thẻ **"Tài khoản Apple ID"** (`Apple ID Accounts`).
2. Nhập Email Apple ID và Mật khẩu.
3. Nhấn **"Đăng nhập"** (`Log In`).
4. Nếu tài khoản có bật Xác thực 2 bước (2FA), một cửa sổ pop-up sẽ hiện lên. Kiểm tra mã xác thực 6 chữ số gửi về thiết bị Apple tin cậy của bạn và nhập vào rồi nhấn Xác nhận.
5. Sau khi đăng nhập thành công, thẻ tài khoản sẽ hiển thị với huy hiệu màu vàng: `Chứng chỉ: Chưa tạo` (`Credentials: Not Minted`).

> [!TIP]
> Bạn có thể đăng nhập nhiều tài khoản Apple ID thuộc các khu vực quốc gia khác nhau (Việt Nam, Mỹ, Nhật Bản, v.v.) và chuyển đổi qua lại bất kỳ lúc nào chỉ với 1 cú click vào nút "Kích hoạt".

---

### Bước 2: Tạo chứng chỉ FairPlay 1 lần duy nhất (Mint Credential)
1. Cắm thiết bị iOS đã Jailbreak vào máy tính qua cáp USB (khuyên dùng) hoặc kết nối chung mạng Wi-Fi với máy tính.
2. Chuyển sang thẻ **"Thiết bị & Mint"** (`Device & Mint`).
3. Thiết lập thông tin kết nối:
   - **Kết nối USB**: Tích chọn ô "Kết nối qua cáp USB (sử dụng iproxy)" *(Khuyên dùng)*.
   - **Kết nối Wi-Fi**: Nhập địa chỉ IP của iPhone (ví dụ: `192.168.1.100`), Cổng `22`, Tài khoản `root`, Mật khẩu `alpine`.
4. Nhấn nút **"Lưu cấu hình"**, sau đó nhấn **"Kiểm tra kết nối"**.
   - Hệ thống sẽ tự động kiểm tra quyền SSH root, AppSync Unified, appinst và tính năng nạp helper.
5. Nhấn nút **"Tạo chứng chỉ FairPlay (Mint Credential)"**.
6. Chờ trong khoảng 5 - 15 giây để hệ thống ký, đẩy payload lên iPhone và trích xuất khóa giải mã.
7. Sau khi hoàn tất, huy hiệu tài khoản chuyển sang màu xanh lá: `Chứng chỉ: Đã tạo` (`Credentials: Minted`)! Kể từ thời điểm này, bạn có thể rút cáp iPhone ra và giải mã thoải mái ngoại tuyến.

---

### Bước 3: Giải mã ứng dụng

#### Cách 1: Tải và giải mã trực tiếp từ App Store
1. Chuyển sang thẻ **"Giải mã IPA"** (`Decrypt IPA`).
2. Chọn nguồn **"Tải từ App Store"**.
3. Nhập **Bundle ID** (ví dụ `com.spotify.client`), **App ID** (ví dụ `324684580`), hoặc đường dẫn URL đầy đủ trên App Store.
4. Cấu hình các tùy chọn nâng cao nếu cần (phiên bản, storefront quốc gia, mở rộng hỗ trợ thiết bị,...).
5. Nhấn **"Bắt đầu giải mã FairPlay"**.
6. Theo dõi tiến trình tải và giải mã trực tiếp trong khung giám sát.

#### Cách 2: Tìm kiếm trên App Store trước khi giải mã
1. Chuyển sang thẻ **"Tìm kiếm App Store"** (`App Store Search`).
2. Gõ tên ứng dụng hoặc từ khóa, chọn quốc gia (VN, US, JP,...).
3. Nhấn nút **"Tải & Giải mã"** trên thẻ thông tin ứng dụng để tự động điền sang màn hình giải mã.

#### Cách 3: Giải mã file IPA có sẵn trên máy tính
1. Tại thẻ **"Giải mã IPA"**, chọn nguồn **"File IPA có sẵn"**.
2. Kéo thả file `.ipa` đã mã hóa vào khung hoặc nhấn vào để duyệt file từ máy tính.
3. Nhấn **"Bắt đầu giải mã FairPlay"**.

---

### Bước 4: Tải bản cũ (Lịch sử phiên bản App Store)
1. Chuyển sang thẻ **"Lịch sử phiên bản"** (`Version History`).
2. Nhập Bundle ID hoặc App ID rồi nhấn **"Xem phiên bản"**.
3. Danh sách toàn bộ các bản phát hành trong quá khứ của ứng dụng từ máy chủ Apple sẽ hiển thị.
4. Nhấn vào nút **"Giải mã bản này"** ở bất kỳ dòng phiên bản cũ nào để tải về và giải mã bản đó!

---

### Bước 5: Quản lý file đã giải mã & Cài đặt vào máy
1. Chuyển sang thẻ **"Tệp đã giải mã"** (`Decrypted Files`).
2. Tại đây bạn có thể:
   - Nhấn **"Mở thư mục trong Explorer"** để truy cập ngay thư mục `./decrypted`.
   - Nhấn **"Tải về máy"** để tải file IPA qua trình duyệt.
   - Nhấn **"Cài đặt vào máy"**: Cắm bất kỳ thiết bị iPhone/iPad nào qua cổng USB, chọn máy trong danh sách thả xuống và cài đặt trực tiếp file IPA đã giải mã vào máy thông qua tiện ích `ideviceinstaller` đi kèm!

---

## ⚙️ Bảng giải thích chi tiết các tùy chọn giải mã nâng cao

| Tùy chọn trên WebGUI | Cờ CLI tương ứng | Ý nghĩa & Tác dụng | Khi nào nên sử dụng |
| :--- | :--- | :--- | :--- |
| **Chế độ giải mã** | `--on-device` / `--no-on-device` | Quyết định giải mã offline bằng khóa FairPlay hay đẩy lên máy Jailbreak để dump. | Nên để **Tự động**. Chỉ bật `--on-device` nếu app chứa phân vùng 32-bit cũ. |
| **Mở rộng UIDeviceFamily** | `--patch-device-type` | Sửa giá trị `UIDeviceFamily` trong file `Info.plist` để mở rộng danh sách thiết bị hỗ trợ. | Dùng khi muốn cài một app dành riêng cho iPad lên iPhone hoặc ngược lại. |
| **Bỏ qua tiện ích mở rộng** | `--skip-appex` | Không giải mã các plugin trong thư mục `Payload/<App>.app/PlugIns/*.appex`. | Giúp tăng tốc độ giải mã khi bạn chỉ cần phân tích binary chính của app. |
| **Kiểm tra nâng cao (Extra Verify)** | `--extra-verify` | So sánh từng byte dữ liệu bên ngoài vùng mã hóa với file gốc để phát hiện sai lệch. | Khuyên dùng khi cần tính toàn vẹn tuyệt đối phục vụ nghiên cứu bảo mật. |
| **Bỏ qua kiểm tra (No Verify)** | `--no-verify` | Bỏ qua bước kiểm tra cấu trúc header Mach-O `cryptid`. | Chỉ dùng khi cấu trúc header của app gặp bất thường đặc biệt. |
| **Phiên bản ID cụ thể** | `--external-version-id` | Ghim tải ứng dụng theo ID phiên bản nội bộ cụ thể trên máy chủ Apple. | Dùng khi muốn rollback về một phiên bản cũ của ứng dụng. |
| **Storefront quốc gia** | `--storefront` | Ghi đè mã vùng quốc gia (ví dụ: `US`, `VN`, `143441`). | Dùng khi tải các ứng dụng bị giới hạn khu vực địa lý. |

---

## 💻 Sử dụng qua Giao diện Dòng lệnh (CLI)

Ngoài WebGUI, `ipadecrypt` hỗ trợ đầy đủ các lệnh qua terminal:

```sh
# Khởi chạy máy chủ Web GUI
ipadecrypt gui --port 54322 --host 127.0.0.1 --output-dir ./decrypted

# Thiết lập tài khoản và tạo chứng chỉ ban đầu qua dòng lệnh
ipadecrypt bootstrap

# Tải và giải mã theo Bundle ID, App ID hoặc liên kết App Store
ipadecrypt decrypt com.spotify.client
ipadecrypt decrypt 324684580 --output ./Spotify.ipa
ipadecrypt decrypt "https://apps.apple.com/vn/app/spotify-music-and-podcasts/id324684580"

# Giải mã một file .ipa đã có sẵn trên máy tính
ipadecrypt decrypt D:\AppStore\encrypted.ipa --output D:\AppStore\decrypted.ipa

# Tải và giải mã một phiên bản cũ cụ thể
ipadecrypt decrypt com.spotify.client --external-version-id 858584852

# Duyệt lịch sử phiên bản trực quan bằng bảng tương tác TUI
ipadecrypt versions com.spotify.client
```

---

## 📁 Cấu trúc thư mục dự án

```
ipadecrypt-webGUI/
├── cmd/
│   └── ipadecrypt/
│       ├── gui.go                 # Máy chủ Web, định tuyến & SPA handler
│       ├── gui_accounts.go        # Quản lý đa tài khoản Apple ID & 2FA
│       ├── gui_batch.go           # Xử lý hàng đợi giải mã hàng loạt
│       ├── gui_decrypt.go         # Điều phối quy trình tải và giải mã
│       ├── gui_device.go          # Giao tiếp SSH, iproxy & mint chứng chỉ
│       ├── gui_files.go           # API quản lý danh sách file giải mã
│       ├── gui_install.go         # Cầu nối cài đặt app vào iOS qua libimobiledevice
│       ├── gui_search.go          # API tìm kiếm ứng dụng App Store
│       ├── gui_versions.go        # API tra cứu lịch sử phiên bản
│       ├── main.go                # Điểm khởi động Cobra CLI
│       └── gui_assets/            # Giao diện Web tĩnh nhúng vào nhị phân
│           ├── index.html         # Mã nguồn HTML Single Page App
│           ├── style.css          # Giao diện CSS Dark/Light responsive
│           └── app.js             # Logic điều khiển giao diện & đa ngôn ngữ
├── helper/                        # Bộ giải mã C hỗ trợ chạy trực tiếp trên iOS (arm64)
├── internal/
│   ├── apphttp/                   # Trình giao tiếp API App Store
│   ├── credmint/                  # Bộ công cụ trích xuất khóa phần cứng FairPlay
│   ├── decrypt/                   # Động cơ giải mã FairPlay Mach-O ngoại tuyến
│   ├── device/                    # Lớp điều khiển SSH, SCP & đường hầm iproxy
│   └── store/                     # Trình giả lập AppleMediaServices attestation
├── tools/                         # Các công cụ libimobiledevice cho Windows tích hợp sẵn
│   ├── ideviceinstaller.exe
│   ├── idevice_id.exe
│   ├── iproxy.exe
│   └── *.dll
├── decrypted/                     # Thư mục mặc định chứa các file IPA đã giải mã
├── build.bat                      # File script biên dịch 1 cú click trên Windows
├── ipadecrypt-gui.bat             # File script chạy nhanh WebGUI (CMD)
├── ipadecrypt-gui.ps1             # File script chạy nhanh WebGUI (PowerShell)
├── BUILDING.md                    # Tài liệu hướng dẫn biên dịch
├── WINDOWS_GUI_GUIDE.md           # Hướng dẫn sử dụng giao diện Windows
├── README.md                      # Tài liệu tiếng Anh
└── README.vi.md                   # Tài liệu tiếng Việt
```

---

## ❓ Xử lý sự cố & Câu hỏi thường gặp (FAQ)

### 1. Báo lỗi `[ERROR] Port 54322 is already in use!`
- Hiện đang có một tiến trình `ipadecrypt.exe` khác đang chạy ngầm trên máy.
- Mở Task Manager trên Windows, tìm tiến trình `ipadecrypt.exe` và tắt đi, hoặc chạy lệnh sau trong PowerShell:
  ```powershell
  Stop-Process -Name ipadecrypt -Force
  ```

### 2. Kiểm tra kết nối SSH đến iPhone bị thất bại
- Đảm bảo trên thiết bị Jailbreak đã cài đặt và khởi chạy gói **OpenSSH**.
- Nếu dùng cáp **USB**: Đảm bảo đã cài đặt iTunes hoặc Apple Devices trên máy tính để Windows nhận driver cổng USB của iPhone. Kiểm tra phần mềm diệt virus không chặn file `tools/iproxy.exe`.
- Nếu dùng mạng **Wi-Fi**: Đảm bảo máy tính và iPhone kết nối chung mạng Wi-Fi cục bộ và không có tường lửa chặn cổng 22.
- Tài khoản SSH mặc định là `root`, mật khẩu là `alpine`.

### 3. Báo lỗi "Cannot download app: license not found"
- Đối với các ứng dụng miễn phí, tài khoản Apple ID của bạn phải từng nhấn "Nhận" (Get) ứng dụng đó ít nhất một lần để Apple cấp quyền sở hữu (license) vào tài khoản. Hãy mở App Store trên một thiết bị bất kỳ có đăng nhập tài khoản đó, nhấn "Nhận" app rồi tiến hành giải mã lại trên ipadecrypt.

### 4. Tại sao sau khi Mint chứng chỉ lại không cần cắm iPhone nữa?
- FairPlay sử dụng khóa phần cứng mã hóa các trang nhị phân của ứng dụng. Trong bước **Mint chứng chỉ**, ipadecrypt thực thi một đoạn code nhỏ trên máy Jailbreak để trích xuất khóa giải mã phần cứng đại diện cho tài khoản Apple ID của bạn. Khi khóa này được lưu trữ an toàn trong máy tính (`~/.ipadecrypt/`), công cụ có thể tính toán và giải mã hoàn toàn ngoại tuyến trên máy tính cho mọi app `arm64` của tài khoản đó!

---

## 🔒 Tuyên bố miễn trừ trách nhiệm pháp lý

Phần mềm này được tạo ra và phát triển nghiêm ngặt chỉ nhằm mục đích nghiên cứu bảo mật thông tin, học tập, nghiên cứu khả năng tương tác hệ thống và lưu trữ cá nhân theo quy định của pháp luật về dịch ngược phần mềm hợp pháp.

Người dùng tuyệt đối không sử dụng công cụ này vào mục đích phát tán trái phép các phần mềm thương mại có bản quyền hoặc vi phạm quyền sở hữu trí tuệ. Tác giả không chịu bất kỳ trách nhiệm pháp lý nào đối với những hành vi sử dụng sai mục đích.

---

## 💖 Lời cảm ơn & Nguồn tham khảo

- [ipadecrypt](https://github.com/londek/ipadecrypt) bởi tác giả **londek** — Nền tảng cốt lõi của bộ giải mã FairPlay ngoại tuyến.
- [ipatool](https://github.com/majd/ipatool) bởi tác giả **majd** — Cơ chế tải ứng dụng App Store mô phỏng AppleMediaServices.
- [fairplay-decrypt](https://github.com/EthanArbuckle/fairplay-decrypt) bởi tác giả **Ethan Arbuckle** — Công trình nghiên cứu trích xuất chứng chỉ FairPlay.
- [libimobiledevice](https://libimobiledevice.org/) — Thư viện mã nguồn mở giao tiếp thiết bị iOS đa nền tảng.
