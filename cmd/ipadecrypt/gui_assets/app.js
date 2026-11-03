// ipadecrypt-webGUI Client Application (Strict Bilingual VI/EN & High-Density UI)

const i18n = {
  vi: {
    window_title: 'ipadecrypt webGUI — Công cụ giải mã IPA FairPlay',
    app_subtitle: 'Công cụ giải mã IPA FairPlay từ App Store',
    auth_checking: 'Đang kiểm tra tài khoản...',
    auth_no_account: 'Chưa đăng nhập Apple ID',
    theme_toggle_title: 'Chuyển giao diện sáng/tối',
    tab_decrypt: 'Giải mã IPA',
    tab_search: 'Tìm kiếm App Store',
    tab_versions: 'Lịch sử phiên bản',
    tab_device: 'Thiết bị & Mint',
    tab_files: 'Tệp đã giải mã',
    tab_account: 'Tài khoản Apple ID',
    tab_guide: 'Hướng dẫn',
    
    job_decrypting: 'Đang giải mã ứng dụng...',
    btn_cancel: 'Hủy',

    decrypt_title: 'Giải mã FairPlay',
    decrypt_desc: 'Giải mã ứng dụng App Store ngoại tuyến (offline arm64) hoặc giải mã trên thiết bị Jailbreak (cho app 32-bit).',
    lbl_source_type: 'Nguồn IPA:',
    btn_src_store: 'Tải từ App Store',
    btn_src_local: 'File IPA có sẵn',
    lbl_bundle_or_id: 'Bundle ID, App ID hoặc URL App Store:',
    ph_decrypt_target: 'e.g. com.spotify.client hoặc 324684580',
    lbl_version_id: 'Phiên bản ID (--external-version-id):',
    ph_decrypt_version: 'Để trống = lấy bản mới nhất',
    lbl_storefront: 'Storefront Quốc gia (--storefront):',
    ph_decrypt_storefront: 'e.g. 143441 hoặc US, VN (Để trống = mặc định theo tài khoản)',
    drop_zone_prompt: 'Nhấn để chọn file hoặc kéo thả file .ipa vào đây',
    lbl_decrypt_options: 'Tùy chọn giải mã của ipadecrypt:',
    lbl_decrypt_mode: 'Chế độ giải mã:',
    opt_mode_auto: 'Tự động (Offline FairPlay arm64, fallback máy Jailbreak nếu 32-bit)',
    opt_mode_on_device: 'Bắt buộc giải mã trên máy Jailbreak (--on-device)',
    opt_mode_no_on_device: 'Chỉ giải mã offline, không dùng máy (--no-on-device)',
    opt_patch_device_type: 'Mở rộng UIDeviceFamily để app cài được trên máy khác loại (--patch-device-type)',
    opt_skip_appex: 'Bỏ qua tiện ích mở rộng PlugIns/*.appex (--skip-appex)',
    opt_extra_verify: 'Kiểm tra Mach-O nâng cao byte-compare với file gốc (--extra-verify)',
    opt_no_verify: 'Bỏ qua bước kiểm tra Mach-O cryptid (--no-verify)',
    btn_start_decrypt: 'Bắt đầu giải mã FairPlay',
    lbl_job_monitor: 'Giám sát tiến trình giải mã',
    lbl_job_monitor_desc: 'Thông tin chi tiết các bước xử lý và nhật ký theo thời gian thực.',
    decrypt_empty_log: 'Chưa có tác vụ giải mã nào đang chạy. Nhập thông tin bên trái và nhấn "Bắt đầu giải mã FairPlay".',
    decrypt_success_title: 'Giải mã hoàn tất thành công!',
    btn_open_folder: 'Mở thư mục lưu file',
    btn_download_ipa: 'Tải file IPA về máy',

    search_title: 'Tìm kiếm trên App Store',
    search_desc: 'Tra cứu thông tin ứng dụng để giải mã hoặc xem danh sách các bản phát hành lịch sử.',
    ph_search_input: 'Tìm tên ứng dụng, từ khóa hoặc Bundle ID...',
    btn_search: 'Tìm kiếm',
    search_loading: 'Đang tìm kiếm trên App Store...',
    search_welcome_title: 'Tra cứu ứng dụng App Store',
    search_welcome_desc: 'Nhập tên hoặc Bundle ID vào ô tìm kiếm để tra cứu nhanh.',
    search_no_results: 'Không tìm thấy ứng dụng nào phù hợp.',
    btn_decrypt: 'Giải mã',
    btn_versions: 'Xem phiên bản',

    versions_title: 'Lịch sử phiên bản App Store',
    versions_desc: 'Xem toàn bộ các bản phát hành lịch sử của ứng dụng và giải mã bản cũ tùy ý (lệnh ipadecrypt versions).',
    ph_versions_target: 'Nhập Bundle ID (e.g. com.spotify.client) hoặc App ID...',
    btn_check_versions: 'Xem phiên bản',
    versions_loading: 'Đang tải danh sách phiên bản từ App Store...',
    th_version_id: 'Version ID',
    th_display_version: 'Phiên bản hiển thị',
    th_release_date: 'Ngày phát hành',
    th_action: 'Thao tác',
    badge_latest: 'Mới nhất',
    btn_view_meta: 'Xem',
    versions_unit: 'phiên bản',

    device_title: 'Cấu hình thiết bị Jailbreak',
    device_desc: 'Dùng để tạo chứng chỉ FairPlay 1 lần cho tài khoản hoặc giải mã các app 32-bit (ipadecrypt bootstrap).',
    lbl_device_usb: 'Kết nối qua cáp USB (sử dụng iproxy)',
    lbl_device_host: 'Địa chỉ IP thiết bị (Wi-Fi):',
    lbl_device_port: 'Cổng SSH (Mặc định 22):',
    lbl_device_user: 'Tài khoản SSH (Mặc định root):',
    lbl_device_password: 'Mật khẩu SSH (Mặc định alpine):',
    lbl_device_udid: 'UDID thiết bị (Tùy chọn khi dùng USB):',
    ph_device_udid: 'Tùy chọn nếu cắm nhiều thiết bị',
    btn_save_config: 'Lưu cấu hình',
    btn_test_connection: 'Kiểm tra kết nối',
    mint_title: 'Chứng chỉ FairPlay & Kiểm tra',
    mint_desc: 'Tạo chứng chỉ 16-byte duy nhất 1 lần để giải mã ngoại tuyến vĩnh viễn.',
    lbl_target_account: 'Tài khoản mục tiêu:',
    btn_mint_now: 'Tạo chứng chỉ FairPlay (Mint Credential)',
    lbl_connection_check: 'Kết quả kiểm tra thiết bị:',
    check_ssh_connecting: 'Đang kết nối SSH đến thiết bị Jailbreak...',
    badge_ready: 'Sẵn sàng',
    badge_missing: 'Chưa có',

    files_title: 'Tệp IPA đã giải mã',
    btn_refresh: 'Làm mới',
    files_empty: 'Chưa có file IPA giải mã nào trong thư mục.',
    th_file_name: 'Tên tệp',
    th_file_size: 'Kích thước',
    th_date: 'Ngày tạo',
    btn_download: 'Tải về',
    btn_delete: 'Xóa',

    accounts_list_title: 'Danh sách tài khoản Apple ID',
    accounts_list_desc: 'Hỗ trợ nhiều tài khoản độc lập. Nhấn "Kích hoạt" để chuyển tài khoản sử dụng.',
    no_accounts_msg: 'Chưa có tài khoản nào được lưu. Đăng nhập bên phải để thêm tài khoản.',
    active_pill: 'Đang sử dụng',
    cred_ok: 'Đã có chứng chỉ FairPlay 16-byte',
    cred_missing: 'Chưa tạo chứng chỉ',
    btn_mint_quick: 'Tạo ngay',
    btn_activate: 'Kích hoạt',
    btn_remove: 'Xóa',
    add_account_title: 'Thêm tài khoản Apple ID',
    add_account_desc: 'Đăng nhập Apple ID để tải ứng dụng và tạo chứng chỉ giải mã.',
    lbl_apple_email: 'Email Apple ID:',
    lbl_apple_password: 'Mật khẩu Apple ID:',
    btn_login: 'Đăng nhập',
    btn_logging_in: 'Đang xác thực với Apple...',

    guide_title: 'Hướng dẫn sử dụng ipadecrypt',
    guide_desc: 'Tìm hiểu cách ipadecrypt giải mã FairPlay và cách thiết lập.',

    modal_2fa_title: 'Xác thực hai yếu tố (2FA)',
    modal_2fa_desc: 'Apple đã gửi mã xác thực 6 chữ số đến thiết bị tin cậy của bạn. Vui lòng nhập mã vào bên dưới:',
    btn_verify: 'Xác nhận',
    modal_mint_title: 'Đang tạo chứng chỉ FairPlay',
    modal_mint_desc: 'Đang kết nối đến thiết bị Jailbreak qua SSH để trích xuất khóa 16-byte...',

    modal_help_title: 'Chi tiết tùy chọn giải mã',
    lbl_help_flag: 'Tham số CLI:',
    lbl_help_effect: 'Tác dụng & Ý nghĩa:',
    lbl_help_usage: 'Khi nào nên dùng:',
    lbl_help_notes: 'Lưu ý:',
    btn_close: 'Đóng'
  },
  en: {
    window_title: 'ipadecrypt webGUI — FairPlay App Store Decrypter',
    app_subtitle: 'FairPlay DRM IPA Decrypter for App Store apps',
    auth_checking: 'Checking accounts...',
    auth_no_account: 'Apple ID Sign-in Required',
    theme_toggle_title: 'Toggle light/dark theme',
    tab_decrypt: 'FairPlay Decrypt',
    tab_search: 'App Store Search',
    tab_versions: 'Version History',
    tab_device: 'Device & Mint',
    tab_files: 'Decrypted IPAs',
    tab_account: 'Apple ID Accounts',
    tab_guide: 'User Guide',

    job_decrypting: 'Decrypting application...',
    btn_cancel: 'Cancel',

    decrypt_title: 'FairPlay Decryption',
    decrypt_desc: 'Decrypt App Store apps offline (arm64) or on a jailbroken iOS device (for 32-bit legacy apps).',
    lbl_source_type: 'IPA Source:',
    btn_src_store: 'App Store Download',
    btn_src_local: 'Local IPA File',
    lbl_bundle_or_id: 'Bundle ID, App ID, or App Store URL:',
    ph_decrypt_target: 'e.g. com.spotify.client or 324684580',
    lbl_version_id: 'Version ID (--external-version-id):',
    ph_decrypt_version: 'Leave blank = latest version',
    lbl_storefront: 'Storefront Country (--storefront):',
    ph_decrypt_storefront: 'e.g. 143441 or US, VN (Leave blank = account default)',
    drop_zone_prompt: 'Click to select or drag and drop .ipa file here',
    lbl_decrypt_options: 'ipadecrypt Decryption Options:',
    lbl_decrypt_mode: 'Decryption Mode:',
    opt_mode_auto: 'Automatic (Offline FairPlay arm64, Jailbreak fallback for 32-bit)',
    opt_mode_on_device: 'Force on-device decryption via Jailbreak (--on-device)',
    opt_mode_no_on_device: 'Offline only, no device fallback (--no-on-device)',
    opt_patch_device_type: 'Patch UIDeviceFamily to run on iPad/Apple Silicon (--patch-device-type)',
    opt_skip_appex: 'Skip app extensions PlugIns/*.appex (--skip-appex)',
    opt_extra_verify: 'Extra Mach-O byte verification against original binary (--extra-verify)',
    opt_no_verify: 'Skip Mach-O cryptid verification (--no-verify)',
    btn_start_decrypt: 'Start FairPlay Decryption',
    lbl_job_monitor: 'Decryption Progress Monitor',
    lbl_job_monitor_desc: 'Real-time logs, stages, and output verification.',
    decrypt_empty_log: 'No active decryption job. Configure options on the left and click "Start FairPlay Decryption".',
    decrypt_success_title: 'Decryption completed successfully!',
    btn_open_folder: 'Open Output Folder',
    btn_download_ipa: 'Download IPA File',

    search_title: 'Search on App Store',
    search_desc: 'Lookup app metadata to decrypt or browse release history.',
    ph_search_input: 'Search app name, keyword, or Bundle ID...',
    btn_search: 'Search',
    search_loading: 'Searching on App Store...',
    search_welcome_title: 'Lookup App Store Apps',
    search_welcome_desc: 'Enter an app name or Bundle ID above to search.',
    search_no_results: 'No apps found matching your query.',
    btn_decrypt: 'Decrypt',
    btn_versions: 'Versions',

    versions_title: 'App Store Version History',
    versions_desc: 'Browse all historical releases and decrypt older versions on-demand (ipadecrypt versions).',
    ph_versions_target: 'Enter Bundle ID (e.g. com.spotify.client) or App ID...',
    btn_check_versions: 'Lookup Versions',
    versions_loading: 'Fetching version history from App Store...',
    th_version_id: 'Version ID',
    th_display_version: 'Display Version',
    th_release_date: 'Release Date',
    th_action: 'Action',
    badge_latest: 'Latest',
    btn_view_meta: 'View',
    versions_unit: 'versions',

    device_title: 'Jailbroken Device Configuration',
    device_desc: 'Used once per account to mint FairPlay credential, or for 32-bit app fallback (ipadecrypt bootstrap).',
    lbl_device_usb: 'Connect via USB cable (using iproxy)',
    lbl_device_host: 'Device IP Address (Wi-Fi):',
    lbl_device_port: 'SSH Port (Default 22):',
    lbl_device_user: 'SSH Username (Default root):',
    lbl_device_password: 'SSH Password (Default alpine):',
    lbl_device_udid: 'Device UDID (Optional for multi-device USB):',
    ph_device_udid: 'Optional if multiple devices connected',
    btn_save_config: 'Save Configuration',
    btn_test_connection: 'Test Connection',
    mint_title: 'FairPlay Credential & Verification',
    mint_desc: 'Extract 16-byte FairPlay key once per account for permanent offline decryption.',
    lbl_target_account: 'Target Account:',
    btn_mint_now: 'Mint FairPlay Credential',
    lbl_connection_check: 'Device Verification Results:',
    check_ssh_connecting: 'Connecting to jailbroken device via SSH...',
    badge_ready: 'Ready',
    badge_missing: 'Missing',

    files_title: 'Decrypted IPA Files',
    btn_refresh: 'Refresh',
    files_empty: 'No decrypted IPA files found in the output folder.',
    th_file_name: 'File Name',
    th_file_size: 'File Size',
    th_date: 'Date Created',
    btn_download: 'Download',
    btn_delete: 'Delete',

    accounts_list_title: 'Apple ID Accounts',
    accounts_list_desc: 'Manage multiple isolated accounts. Click "Activate" to switch active session.',
    no_accounts_msg: 'No accounts added yet. Sign in on the right to add an account.',
    active_pill: 'Active',
    cred_ok: '16-byte Credential Minted',
    cred_missing: 'No Credential Minted',
    btn_mint_quick: 'Mint Now',
    btn_activate: 'Activate',
    btn_remove: 'Remove',
    add_account_title: 'Add Apple ID Account',
    add_account_desc: 'Sign in with Apple ID to purchase/download apps and mint credentials.',
    lbl_apple_email: 'Apple ID Email:',
    lbl_apple_password: 'Apple ID Password:',
    btn_login: 'Sign In',
    btn_logging_in: 'Authenticating with Apple...',

    guide_title: 'ipadecrypt Guide',
    guide_desc: 'Learn how FairPlay decryption works and how to configure your environment.',

    modal_2fa_title: 'Two-Factor Authentication (2FA)',
    modal_2fa_desc: 'Apple sent a 6-digit verification code to your trusted devices. Enter it below:',
    btn_verify: 'Confirm',
    modal_mint_title: 'Minting FairPlay Credential',
    modal_mint_desc: 'Connecting to jailbroken device via SSH to extract 16-byte key...',

    modal_help_title: 'Decryption Option Details',
    lbl_help_flag: 'CLI Flag:',
    lbl_help_effect: 'Description & Effect:',
    lbl_help_usage: 'Recommended Usage:',
    lbl_help_notes: 'Important Notes:',
    btn_close: 'Close'
  }
};

// Rich Help Information Data Dictionary for (?) popups
const optionHelpData = {
  decrypt_mode: {
    flag: '--on-device / --no-on-device',
    title: {
      vi: 'Chế độ giải mã (Decryption Mode)',
      en: 'Decryption Mode'
    },
    effect: {
      vi: 'Xác định môi trường thực thi giải mã nhị phân Mach-O:\n• Tự động (Auto): Mặc định ipadecrypt giải mã ngoại tuyến (offline) trên PC cho kiến trúc 64-bit (arm64) bằng khóa 16-byte đã mint. Nếu phát hiện ứng dụng cũ chứa mã 32-bit (armv7/armv7s), công cụ sẽ tự động kết nối iPhone/iPad Jailbreak để giải mã trên máy.\n• Bắt buộc trên máy (--on-device): Bỏ qua giải mã offline, luôn đẩy ứng dụng vào máy Jailbreak và dump bộ nhớ decrypted.\n• Chỉ offline (--no-on-device): Chỉ giải mã trên PC, báo lỗi nếu gặp ứng dụng 32-bit thay vì kết nối thiết bị.',
      en: 'Controls where and how Mach-O FairPlay decryption occurs:\n• Automatic (Auto): Defaults to PC-offline decryption for modern 64-bit (arm64) architectures using your 16-byte credential. Automatically falls back to on-device decryption via SSH/USB on a jailbroken iOS device if a 32-bit (armv7/armv7s) binary is detected.\n• Force On-Device (--on-device): Bypasses offline PC decryption and forces dumping unencrypted binary memory directly on the jailbroken device.\n• Offline Only (--no-on-device): Decrypts strictly on PC; raises an error if 32-bit code is encountered instead of attempting device connection.'
    },
    usage: {
      vi: 'Nên để "Tự động" cho mọi trường hợp thông thường. Chỉ chọn "--on-device" khi gặp app có bảo vệ chống can thiệp đặc thù hoặc nhị phân bị lỗi giải mã offline.',
      en: 'Leave on "Automatic" for 99% of tasks. Use "--on-device" if an app has anti-tamper mechanisms that fail offline decryption.'
    },
    notes: {
      vi: 'Giải mã offline yêu cầu tài khoản đã Mint chứng chỉ FairPlay 16-byte thành công ở tab "Thiết bị & Mint".',
      en: 'Offline decryption requires that your Apple ID has already minted its 16-byte FairPlay credential under the "Device & Mint" tab.'
    }
  },
  patch_device_type: {
    flag: '--patch-device-type',
    title: {
      vi: 'Mở rộng UIDeviceFamily (--patch-device-type)',
      en: 'Patch UIDeviceFamily (--patch-device-type)'
    },
    effect: {
      vi: 'Trong iOS, trường UIDeviceFamily trong tệp Info.plist và Mach-O quy định loại phần cứng được phép chạy (1 = iPhone/iPod, 2 = iPad). Khi bật tùy chọn này, ipadecrypt sẽ tự động sửa đổi giá trị thành [1, 2].',
      en: 'Apple uses UIDeviceFamily in Info.plist and Mach-O headers to constrain apps to iPhone (1) or iPad (2). Enabling this flag patches UIDeviceFamily to [1, 2].'
    },
    usage: {
      vi: 'Rất hữu ích khi bạn muốn cài các ứng dụng vốn chỉ dành cho iPhone (như Instagram, WhatsApp...) lên iPad hoặc máy tính Mac chip M1/M2/M3/M4 để hiển thị giao diện toàn màn hình thay vì khung điện thoại thu nhỏ.',
      en: 'Essential when sideloading iPhone-only apps (e.g. Instagram, WhatsApp) onto iPads or Apple Silicon Macs so they run fullscreen without phone letterboxing.'
    },
    notes: {
      vi: 'Một số ứng dụng có layout cố định cho tỉ lệ màn hình iPhone có thể xuất hiện khoảng đen ở hai bên viền iPad.',
      en: 'Some apps hardcode phone aspect ratios and may render black sidebars on tablet screens.'
    }
  },
  skip_appex: {
    flag: '--skip-appex',
    title: {
      vi: 'Bỏ qua PlugIns/*.appex (--skip-appex)',
      en: 'Skip App Extensions (--skip-appex)'
    },
    effect: {
      vi: 'Các ứng dụng hiện đại thường đi kèm nhiều tiện ích mở rộng trong thư mục PlugIns/ (Widget màn hình, Share Extension, Notification Service, Siri Intents...). Mỗi extension là một nhị phân Mach-O riêng biệt bị mã hóa DRM. Cờ này sẽ bỏ qua việc giải mã các extension này, chỉ giải mã tệp nhị phân chính của ứng dụng.',
      en: 'Modern iOS apps contain numerous app extensions under PlugIns/ (Home Widgets, Share Extensions, Notification Services, Siri Intents). Each extension is an independently encrypted Mach-O binary. This flag skips decrypting them and only decrypts the main application binary.'
    },
    usage: {
      vi: 'Bật khi bạn cần tăng tốc độ giải mã nhanh gấp 2-3 lần đối với các app nặng (Facebook, TikTok, Shopee...) và chỉ quan tâm đến việc chạy ứng dụng chính.',
      en: 'Enable to accelerate decryption speed by 2x-3x on large apps (Facebook, TikTok, etc.) when you only need the main app functionality.'
    },
    notes: {
      vi: 'Nếu bạn cần sử dụng Widget hoặc Notification Extension của app sau khi cài đặt, không nên bật cờ này.',
      en: 'If you need iOS Home Screen Widgets or Rich Notification extensions to function, do not enable this flag.'
    }
  },
  extra_verify: {
    flag: '--extra-verify',
    title: {
      vi: 'Kiểm tra Mach-O nâng cao (--extra-verify)',
      en: 'Extra Mach-O Byte Verification (--extra-verify)'
    },
    effect: {
      vi: 'Sau khi giải mã, ipadecrypt sẽ thực hiện so sánh từng byte (byte-level comparison) của toàn bộ các vùng không mã hóa giữa tệp Mach-O đã giải mã và tệp gốc từ App Store.',
      en: 'Performs an exhaustive byte-level diff between the decrypted Mach-O binary and the original App Store binary across all non-encrypted slices, load commands, and data segments.'
    },
    usage: {
      vi: 'Đảm bảo 100% không xảy ra lỗi sai lệch byte trong quá trình bóc tách DRM. Khuyên dùng cho các chuyên gia dịch ngược (Reverse Engineering), phân tích mã độc hoặc nghiên cứu bảo mật iOS.',
      en: 'Guarantees absolute binary integrity with zero corrupted bytes. Highly recommended for reverse engineering, binary analysis, and security research.'
    },
    notes: {
      vi: 'Quá trình kiểm tra tốn thêm khoảng 1-2 giây cho các ứng dụng có nhị phân lớn trên 100MB.',
      en: 'Adds roughly 1-2 seconds of verification time for binaries exceeding 100MB.'
    }
  },
  no_verify: {
    flag: '--no-verify',
    title: {
      vi: 'Bỏ qua kiểm tra Mach-O cryptid (--no-verify)',
      en: 'Skip Cryptid Verification (--no-verify)'
    },
    effect: {
      vi: 'Mặc định sau khi giải mã, ipadecrypt sẽ đọc lại load command LC_ENCRYPTION_INFO để đảm bảo cờ cryptid đã được chuyển thành 0 (đã giải mã thành công). Cờ này tắt hoàn toàn bước xác minh đó.',
      en: 'By default, ipadecrypt inspects LC_ENCRYPTION_INFO(_64) to confirm cryptid == 0 (fully decrypted). This flag disables that verification step entirely.'
    },
    usage: {
      vi: 'Chỉ nên dùng khi giải mã các nhị phân dị biệt, ứng dụng tự chế hoặc bản build thử nghiệm có cấu trúc header không theo chuẩn thông thường của Apple khiến bộ kiểm tra báo lỗi giả.',
      en: 'Only recommended for experimental, modified, or custom binaries with non-standard Apple headers that cause false verification failures.'
    },
    notes: {
      vi: 'Nếu bật cờ này, bạn sẽ không được cảnh báo nếu nhị phân chưa được giải mã hoàn toàn.',
      en: 'Disabling verification means you will not be warned if a binary remains partially encrypted.'
    }
  },
  version_id: {
    flag: '--external-version-id <id>',
    title: {
      vi: 'Phiên bản ID (--external-version-id)',
      en: 'Historical Version ID (--external-version-id)'
    },
    effect: {
      vi: 'Apple gán cho mỗi lần phát hành ứng dụng trên App Store một mã số nhận diện phiên bản độc nhất (external-version-id, ví dụ 845239120). Tham số này cho phép bạn tải và giải mã chính xác một phiên bản cũ cụ thể trong quá khứ.',
      en: 'Apple assigns a unique integer ID to every historical app release on the App Store (e.g. 845239120). This parameter instructs ipadecrypt to download and decrypt that exact historical release instead of the latest one.'
    },
    usage: {
      vi: 'Dùng khi bạn muốn hạ cấp (downgrade) ứng dụng, tìm lại bản cũ không có quảng cáo, hoặc nghiên cứu sự thay đổi giữa các phiên bản. Bạn có thể sang tab "Lịch sử phiên bản" để tra cứu toàn bộ danh sách ID.',
      en: 'Useful for downgrading apps, retrieving older ad-free releases, or conducting differential reverse engineering. Browse all IDs under the "Version History" tab.'
    },
    notes: {
      vi: 'Để trống để luôn tải và giải mã phiên bản mới nhất trên App Store.',
      en: 'Leave empty to automatically download and decrypt the latest public release.'
    }
  },
  storefront: {
    flag: '--storefront <code/country>',
    title: {
      vi: 'Storefront Quốc gia (--storefront)',
      en: 'Storefront Country Code (--storefront)'
    },
    effect: {
      vi: 'Storefront là mã định danh quốc gia của kho ứng dụng Apple (ví dụ: 143441 cho Mỹ / US, 143471 cho Việt Nam / VN, 143462 cho Nhật Bản / JP). Tham số này ép buộc máy chủ kết nối đến kho ứng dụng của quốc gia mong muốn.',
      en: 'Storefront specifies Apple\'s regional App Store catalog (e.g. 143441 for US, 143471 for Vietnam, 143462 for Japan). This parameter forces requests through the chosen regional store.'
    },
    usage: {
      vi: 'Dùng khi tải các ứng dụng bị giới hạn theo vùng địa lý (chỉ có ở Store Mỹ hoặc Nhật). Bạn có thể nhập mã số (143441) hoặc mã 2 chữ cái (US, VN, JP).',
      en: 'Useful when downloading region-locked apps available only in specific countries. You can enter either the numeric ID (143441) or 2-letter country code (US, VN, JP).'
    },
    notes: {
      vi: 'Để trống để mặc định sử dụng Storefront tương ứng với tài khoản Apple ID đang kích hoạt.',
      en: 'Leave blank to use the storefront country of your currently active Apple ID.'
    }
  }
};

// Global Application State
const state = {
  lang: localStorage.getItem('ipadecrypt_lang') === 'en' ? 'en' : 'vi',
  theme: localStorage.getItem('ipadecrypt_theme') === 'light' ? 'light' : 'dark',
  activeTab: 'decrypt',
  sourceMode: 'store', // 'store' or 'local'
  accounts: [],
  activeAccount: null,
  activeJob: null,
  jobPollTimer: null,
  selectedLocalFile: null
};

// --- Initialization ---
document.addEventListener('DOMContentLoaded', () => {
  applyTheme(state.theme);
  applyLanguage(state.lang);
  loadAppStatus();
  loadAccounts();
  loadDeviceConfig();
  loadDecryptedFiles();

  // Setup drag and drop for local IPA file
  const dropZone = document.getElementById('ipa-drop-zone');
  if (dropZone) {
    dropZone.addEventListener('dragover', (e) => {
      e.preventDefault();
      dropZone.classList.add('dragover');
    });
    dropZone.addEventListener('dragleave', () => dropZone.classList.remove('dragover'));
    dropZone.addEventListener('drop', (e) => {
      e.preventDefault();
      dropZone.classList.remove('dragover');
      if (e.dataTransfer.files && e.dataTransfer.files[0]) {
        handleLocalFileSelected(e.dataTransfer.files[0]);
      }
    });
  }

  // Keyboard shortcut ESC to close modals
  document.addEventListener('keydown', (e) => {
    if (e.key === 'Escape') {
      closeOptionHelp();
      close2FAModal();
    }
  });
});

// --- Theme Management ---
function toggleTheme() {
  state.theme = state.theme === 'dark' ? 'light' : 'dark';
  localStorage.setItem('ipadecrypt_theme', state.theme);
  applyTheme(state.theme);
}

function applyTheme(theme) {
  document.documentElement.setAttribute('data-theme', theme);
  const darkIcon = document.getElementById('theme-icon-dark');
  const lightIcon = document.getElementById('theme-icon-light');
  if (darkIcon && lightIcon) {
    darkIcon.style.display = theme === 'dark' ? 'block' : 'none';
    lightIcon.style.display = theme === 'light' ? 'block' : 'none';
  }
}

// --- Bilingual Language Switcher (Strict VI / EN) ---
function toggleLanguage() {
  state.lang = state.lang === 'vi' ? 'en' : 'vi';
  localStorage.setItem('ipadecrypt_lang', state.lang);
  applyLanguage(state.lang);
}

function t(key) {
  const dict = i18n[state.lang] || i18n.vi;
  return dict[key] || i18n.vi[key] || key;
}

function applyLanguage(lang) {
  const dict = i18n[lang] || i18n.vi;

  // Translate all DOM elements with data-i18n
  document.querySelectorAll('[data-i18n]').forEach(el => {
    const key = el.getAttribute('data-i18n');
    if (dict[key]) el.textContent = dict[key];
  });

  // Translate placeholders
  document.querySelectorAll('[data-i18n-placeholder]').forEach(el => {
    const key = el.getAttribute('data-i18n-placeholder');
    if (dict[key]) el.placeholder = dict[key];
  });

  // Translate titles
  document.querySelectorAll('[data-i18n-title]').forEach(el => {
    const key = el.getAttribute('data-i18n-title');
    if (dict[key]) el.title = dict[key];
  });

  // Update language toggle button label
  const langLabel = document.getElementById('lang-label');
  if (langLabel) langLabel.textContent = lang.toUpperCase();

  // Render Guide Tab with full language content
  renderGuide();

  // Re-render dynamic elements
  renderAccountPill();
  renderAccountsList();
  renderMintTabAccount();
}

function renderGuide() {
  const guideBody = document.getElementById('guide-body');
  if (!guideBody) return;

  if (state.lang === 'en') {
    guideBody.innerHTML = `
      <h4 class="font-bold text-primary mb-1">1. FairPlay Decryption Mechanism</h4>
      <p class="text-secondary mb-3">ipadecrypt decrypts Apple FairPlay-protected App Store packages directly on your PC (offline) <strong>without needing a tethered iPhone/iPad during normal decryption</strong>.</p>
      <p class="text-secondary mb-3">A jailbroken iOS device is required <strong>only once per Apple ID account</strong> to extract (mint) the 16-byte FairPlay credential stored at <code>~/.ipadecrypt/credentials/&lt;dsid&gt;.bin</code>. Once minted, all modern arm64 apps are decrypted completely offline on your PC.</p>

      <h4 class="font-bold text-primary mb-1">2. Multi-Account Sandboxing</h4>
      <p class="text-secondary mb-3">Every Apple ID has a unique DSID (Directory Services Identifier). ipadecrypt isolates sessions and cookie jars per account in <code>~/.ipadecrypt/accounts/&lt;id&gt;/</code>, preventing any session collisions or cookie corruption when switching accounts.</p>

      <h4 class="font-bold text-primary mb-1">3. Legacy 32-bit (armv7) Support</h4>
      <p class="text-secondary">Legacy 32-bit iOS applications cannot be decrypted offline on x86_64/arm64 PCs. When a 32-bit Mach-O slice is encountered, ipadecrypt automatically invokes the on-device helper fallback via SSH/USB (using <code>AppSync Unified</code> and <code>appinst</code>) to decrypt directly on the jailbroken device.</p>
    `;
  } else {
    guideBody.innerHTML = `
      <h4 class="font-bold text-primary mb-1">1. Cơ chế giải mã FairPlay của ipadecrypt</h4>
      <p class="text-secondary mb-3">ipadecrypt giải mã các ứng dụng App Store được bảo vệ bởi công nghệ FairPlay của Apple trực tiếp trên máy tính của bạn (offline) mà <strong>không cần cắm thiết bị iPhone/iPad liên tục</strong>.</p>
      <p class="text-secondary mb-3">Thiết bị Jailbreak chỉ cần kết nối <strong>duy nhất 1 lần cho mỗi tài khoản Apple ID</strong> để trích xuất (mint) chứng chỉ FairPlay 16-byte lưu tại <code>~/.ipadecrypt/credentials/&lt;dsid&gt;.bin</code>. Sau bước này, mọi ứng dụng arm64 hiện đại đều được giải mã ngoại tuyến trên PC.</p>

      <h4 class="font-bold text-primary mb-1">2. Quản lý nhiều tài khoản (Multi-Account)</h4>
      <p class="text-secondary mb-3">Mỗi Apple ID có một mã nhận dạng DSID riêng biệt. ipadecrypt lưu cookie và chứng chỉ độc lập cho từng tài khoản, không bị ghi đè hay xung đột khi chuyển đổi qua lại. Khi giải mã ứng dụng mua bởi tài khoản nào, bạn chỉ cần chọn tài khoản đó làm tài khoản đang hoạt động (Active).</p>

      <h4 class="font-bold text-primary mb-1">3. Đối với ứng dụng cũ 32-bit (armv7)</h4>
      <p class="text-secondary">Các bản build rất cũ còn chứa slice 32-bit không thể giải mã offline trên máy tính x86_64/arm64. ipadecrypt sẽ tự động kích hoạt cơ chế fallback giải mã trực tiếp trên thiết bị Jailbreak (thông qua <code>AppSync Unified</code> và <code>appinst</code>).</p>
    `;
  }
}

// --- Option Help Modal (?) ---
function openOptionHelp(key) {
  const item = optionHelpData[key];
  if (!item) return;

  const modal = document.getElementById('modal-option-help');
  const titleEl = document.getElementById('help-modal-title');
  const flagEl = document.getElementById('help-modal-flag');
  const effectEl = document.getElementById('help-modal-effect');
  const usageEl = document.getElementById('help-modal-usage');
  const notesEl = document.getElementById('help-modal-notes');

  const lang = state.lang === 'en' ? 'en' : 'vi';

  if (titleEl) titleEl.textContent = item.title[lang] || item.title.vi;
  if (flagEl) flagEl.textContent = item.flag;
  if (effectEl) effectEl.innerHTML = (item.effect[lang] || item.effect.vi).replace(/\n/g, '<br>');
  if (usageEl) usageEl.innerHTML = (item.usage[lang] || item.usage.vi).replace(/\n/g, '<br>');
  if (notesEl) notesEl.innerHTML = (item.notes[lang] || item.notes.vi).replace(/\n/g, '<br>');

  if (modal) modal.style.display = 'flex';
}

function closeOptionHelp() {
  const modal = document.getElementById('modal-option-help');
  if (modal) modal.style.display = 'none';
}

// --- Tab Switching ---
function switchTab(tabName) {
  state.activeTab = tabName;
  document.querySelectorAll('.tab-btn').forEach(btn => {
    btn.classList.toggle('active', btn.getAttribute('data-tab') === tabName);
  });
  document.querySelectorAll('.tab-content').forEach(sec => {
    sec.classList.toggle('active', sec.id === `tab-${tabName}`);
  });

  if (tabName === 'files') loadDecryptedFiles();
  if (tabName === 'account') loadAccounts();
}

// --- Toast Notifications ---
function showToast(message, type = 'info') {
  const container = document.getElementById('toast-container');
  if (!container) return;

  const toast = document.createElement('div');
  toast.className = `toast toast-${type}`;
  toast.textContent = message;
  container.appendChild(toast);

  setTimeout(() => {
    toast.style.opacity = '0';
    setTimeout(() => toast.remove(), 250);
  }, 4000);
}

// --- Source Mode Selector (App Store vs Local IPA) ---
function selectSourceMode(mode) {
  state.sourceMode = mode;
  const btnStore = document.getElementById('btn-src-store');
  const btnLocal = document.getElementById('btn-src-local');
  const storeFields = document.getElementById('src-store-fields');
  const localFields = document.getElementById('src-local-fields');

  if (mode === 'store') {
    btnStore.className = 'btn btn-primary btn-sm flex-1';
    btnLocal.className = 'btn btn-outline btn-sm flex-1';
    storeFields.style.display = 'block';
    localFields.style.display = 'none';
  } else {
    btnStore.className = 'btn btn-outline btn-sm flex-1';
    btnLocal.className = 'btn btn-primary btn-sm flex-1';
    storeFields.style.display = 'none';
    localFields.style.display = 'block';
  }
}

function handleLocalFileSelected(file) {
  if (!file) return;
  state.selectedLocalFile = file;
  const textEl = document.getElementById('drop-zone-text');
  const infoEl = document.getElementById('selected-file-info');

  if (textEl) textEl.textContent = file.name;
  if (infoEl) infoEl.textContent = formatBytes(file.size);
}

// --- Status & Multi-Account Management ---
async function loadAppStatus() {
  try {
    const res = await fetch('/api/status');
    const data = await res.json();
    if (data.success) {
      if (data.version) {
        const vEl = document.getElementById('brand-version');
        if (vEl) {
          vEl.textContent = 'v' + data.version;
          vEl.style.display = 'inline';
        }
      }
      if (data.outputDir) {
        const oEl = document.getElementById('files-output-dir-label');
        if (oEl) oEl.textContent = data.outputDir;
      }
    }
  } catch (err) {
    console.error('Failed to load status:', err);
  }
}

async function loadAccounts() {
  try {
    const res = await fetch('/api/accounts');
    const data = await res.json();
    if (data.success) {
      state.accounts = data.accounts || [];
      state.activeAccount = state.accounts.find(a => a.active) || null;
      renderAccountPill();
      renderAccountsList();
      renderMintTabAccount();
    }
  } catch (err) {
    console.error('Failed to load accounts:', err);
  }
}

function renderAccountPill() {
  const pill = document.getElementById('account-status-pill');
  const text = document.getElementById('account-status-text');
  if (!pill || !text) return;

  if (state.activeAccount) {
    pill.classList.add('active');
    text.textContent = state.activeAccount.email;
    pill.title = state.activeAccount.name ? `${state.activeAccount.name} (${state.activeAccount.storefrontCountry || ''})` : state.activeAccount.email;
  } else {
    pill.classList.remove('active');
    text.textContent = t('auth_no_account');
  }
}

function renderAccountsList() {
  const container = document.getElementById('accounts-cards-list');
  if (!container) return;

  if (state.accounts.length === 0) {
    container.innerHTML = `<p class="text-secondary text-xs p-3">${t('no_accounts_msg')}</p>`;
    return;
  }

  container.innerHTML = state.accounts.map(acc => {
    const isActive = acc.active;
    const country = acc.storefrontCountry ? `(${acc.storefrontCountry})` : '';
    const credBadge = acc.hasCredential 
      ? `<span class="badge-cred-ok"><svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3"><polyline points="20 6 9 17 4 12"></polyline></svg> ${t('cred_ok')}</span>`
      : `<span class="badge-cred-missing"><svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"></circle><line x1="12" y1="8" x2="12" y2="12"></line><line x1="12" y1="16" x2="12.01" y2="16"></line></svg> ${t('cred_missing')}</span>`;

    return `
      <div class="card p-3 ${isActive ? 'border-primary' : ''}">
        <div class="flex-between">
          <div>
            <div class="flex-gap items-center">
              <span class="font-semibold text-primary text-sm">${escapeHtml(acc.email)}</span>
              ${country ? `<span class="badge badge-primary">${country}</span>` : ''}
              ${isActive ? `<span class="badge badge-success font-bold">${t('active_pill')}</span>` : ''}
            </div>
            <div class="text-xs text-secondary mt-1 font-mono">DSID: ${escapeHtml(acc.dsid || 'N/A')}</div>
            <div class="mt-2 flex-gap">
              ${credBadge}
              ${!acc.hasCredential ? `<button class="btn btn-outline btn-xs" onclick="quickMintAccount('${acc.id}')">${t('btn_mint_quick')}</button>` : ''}
            </div>
          </div>
          <div class="flex-gap">
            ${!isActive ? `<button class="btn btn-outline btn-sm" onclick="switchActiveAccount('${acc.id}')">${t('btn_activate')}</button>` : ''}
            <button class="btn btn-danger btn-sm" onclick="removeAccount('${acc.id}')">${t('btn_remove')}</button>
          </div>
        </div>
      </div>
    `;
  }).join('');
}

function renderMintTabAccount() {
  const emailEl = document.getElementById('mint-account-email');
  const badgeEl = document.getElementById('mint-account-badge');
  if (!emailEl || !badgeEl) return;

  if (state.activeAccount) {
    emailEl.textContent = `${state.activeAccount.email} (DSID: ${state.activeAccount.dsid || 'N/A'})`;
    badgeEl.innerHTML = state.activeAccount.hasCredential
      ? `<span class="badge-cred-ok">${t('cred_ok')}</span>`
      : `<span class="badge-cred-missing">${t('cred_missing')}</span>`;
  } else {
    emailEl.textContent = t('auth_no_account');
    badgeEl.innerHTML = '';
  }
}

async function switchActiveAccount(id) {
  try {
    const res = await fetch('/api/accounts/activate', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ id })
    });
    const data = await res.json();
    if (data.success) {
      showToast(state.lang === 'vi' ? 'Đã kích hoạt tài khoản được chọn!' : 'Account activated successfully!', 'success');
      loadAccounts();
    } else {
      showToast(data.error || 'Failed to activate account', 'danger');
    }
  } catch (err) {
    showToast(err.message, 'danger');
  }
}

async function removeAccount(id) {
  const confirmMsg = state.lang === 'vi' ? 'Bạn có chắc chắn muốn xóa tài khoản này khỏi ứng dụng?' : 'Are you sure you want to remove this account?';
  if (!confirm(confirmMsg)) return;

  try {
    const res = await fetch('/api/accounts/remove', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ id })
    });
    const data = await res.json();
    if (data.success) {
      showToast(state.lang === 'vi' ? 'Đã xóa tài khoản' : 'Account removed', 'success');
      loadAccounts();
    } else {
      showToast(data.error || 'Failed to remove account', 'danger');
    }
  } catch (err) {
    showToast(err.message, 'danger');
  }
}

let pendingLoginEmail = '';
let pendingLoginPassword = '';

async function doLogin() {
  const email = document.getElementById('login-email').value.trim();
  const password = document.getElementById('login-password').value;

  if (!email || !password) {
    showToast(state.lang === 'vi' ? 'Vui lòng nhập email và mật khẩu' : 'Please enter Apple ID email and password', 'warning');
    return;
  }

  const btn = document.getElementById('btn-login');
  btn.disabled = true;
  btn.textContent = t('btn_logging_in');

  pendingLoginEmail = email;
  pendingLoginPassword = password;

  try {
    const res = await fetch('/api/auth/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password })
    });
    const data = await res.json();

    if (data.authCodeRequired) {
      open2FAModal();
    } else if (data.success) {
      showToast(state.lang === 'vi' ? 'Đăng nhập thành công!' : 'Sign in successful!', 'success');
      document.getElementById('login-password').value = '';
      loadAccounts();
    } else {
      showToast(data.error || 'Login failed', 'danger');
    }
  } catch (err) {
    showToast(err.message, 'danger');
  } finally {
    btn.disabled = false;
    btn.textContent = t('btn_login');
  }
}

function open2FAModal() {
  const modal = document.getElementById('modal-2fa');
  const input = document.getElementById('input-2fa-code');
  if (modal && input) {
    modal.style.display = 'flex';
    input.value = '';
    setTimeout(() => input.focus(), 150);
  }
}

function close2FAModal() {
  const modal = document.getElementById('modal-2fa');
  if (modal) modal.style.display = 'none';
}

async function submit2FACode() {
  const input = document.getElementById('input-2fa-code');
  const code = input ? input.value.trim() : '';

  if (!code || code.length < 6) {
    showToast(state.lang === 'vi' ? 'Mã 2FA phải có 6 chữ số' : 'Enter 6-digit 2FA code', 'warning');
    return;
  }

  close2FAModal();

  try {
    const res = await fetch('/api/auth/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        email: pendingLoginEmail,
        password: pendingLoginPassword,
        authCode: code
      })
    });
    const data = await res.json();

    if (data.success) {
      showToast(state.lang === 'vi' ? 'Xác thực và đăng nhập thành công!' : 'Authenticated and logged in!', 'success');
      document.getElementById('login-password').value = '';
      loadAccounts();
    } else {
      showToast(data.error || '2FA authentication failed', 'danger');
    }
  } catch (err) {
    showToast(err.message, 'danger');
  }
}

// --- Trigger Decryption (FairPlay Core Engine) ---
async function triggerDecryption() {
  const decryptMode = document.getElementById('decrypt-mode').value;
  const onDevice = decryptMode === 'on-device';
  const noOnDevice = decryptMode === 'no-on-device';
  const patchDeviceType = document.getElementById('opt-patch-device-type').checked;
  const skipAppex = document.getElementById('opt-skip-appex').checked;
  const extraVerify = document.getElementById('opt-extra-verify').checked;
  const noVerify = document.getElementById('opt-no-verify').checked;

  if (state.sourceMode === 'local') {
    if (!state.selectedLocalFile) {
      showToast(state.lang === 'vi' ? 'Vui lòng chọn file .ipa từ máy tính trước' : 'Please select a local .ipa file first', 'warning');
      return;
    }

    const formData = new FormData();
    formData.append('file', state.selectedLocalFile);
    if (state.activeAccount) formData.append('accountId', state.activeAccount.id);

    showToast(state.lang === 'vi' ? 'Đang gửi file IPA lên máy chủ...' : 'Uploading IPA file...', 'info');

    try {
      const res = await fetch('/api/decrypt/local', {
        method: 'POST',
        body: formData
      });
      const data = await res.json();
      if (data.success) {
        showToast(state.lang === 'vi' ? 'Đã bắt đầu giải mã file IPA' : 'IPA decryption started', 'success');
        startJobPolling(data.jobId);
      } else {
        showToast(data.error || 'Failed to start local decrypt', 'danger');
      }
    } catch (err) {
      showToast(err.message, 'danger');
    }
    return;
  }

  // App Store Source
  const target = document.getElementById('decrypt-target').value.trim();
  const versionId = document.getElementById('decrypt-version-id').value.trim();
  const storefront = document.getElementById('decrypt-storefront').value.trim();

  if (!target) {
    showToast(state.lang === 'vi' ? 'Vui lòng nhập Bundle ID hoặc App ID' : 'Please enter Bundle ID or App ID', 'warning');
    return;
  }

  const payload = {
    target,
    versionId,
    storefront,
    onDevice,
    noOnDevice,
    patchDeviceType,
    skipAppex,
    extraVerify,
    noVerify
  };

  if (state.activeAccount) payload.accountId = state.activeAccount.id;

  try {
    const res = await fetch('/api/decrypt', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload)
    });
    const data = await res.json();
    if (data.success) {
      showToast(state.lang === 'vi' ? 'Đã khởi chạy tác vụ giải mã FairPlay!' : 'FairPlay decryption started!', 'success');
      startJobPolling(data.jobId);
    } else {
      showToast(data.error || 'Failed to start decrypt', 'danger');
    }
  } catch (err) {
    showToast(err.message, 'danger');
  }
}

function startJobPolling(jobId) {
  if (state.jobPollTimer) clearInterval(state.jobPollTimer);

  const banner = document.getElementById('active-jobs-banner');
  if (banner) banner.style.display = 'block';

  // Live box in Tab 1
  const emptyLog = document.getElementById('decrypt-empty-log');
  const activeLog = document.getElementById('decrypt-active-log');
  const resultBox = document.getElementById('log-result-box');
  if (emptyLog) emptyLog.style.display = 'none';
  if (activeLog) activeLog.style.display = 'block';
  if (resultBox) resultBox.style.display = 'none';

  state.jobPollTimer = setInterval(async () => {
    try {
      const res = await fetch(`/api/decrypt/status?id=${jobId}`);
      const data = await res.json();

      if (data.success && data.job) {
        const j = data.job;
        state.activeJob = j;
        updateActiveJobBanner(j);
        updateLiveMonitor(j);

        if (j.status === 'completed') {
          clearInterval(state.jobPollTimer);
          showToast(state.lang === 'vi' ? `Hoàn tất giải mã: ${j.outputName || j.bundleId}` : `Decryption completed: ${j.outputName || j.bundleId}`, 'success');
          loadDecryptedFiles();

          if (resultBox) {
            resultBox.style.display = 'block';
            document.getElementById('log-file-name').textContent = `${j.outputName} (${formatBytes(j.outputSize)})`;
            const dlBtn = document.getElementById('log-download-btn');
            if (dlBtn) dlBtn.href = `/api/files/download?name=${encodeURIComponent(j.outputName)}`;
          }

          setTimeout(() => { if (banner) banner.style.display = 'none'; }, 6000);
        } else if (j.status === 'error') {
          clearInterval(state.jobPollTimer);
          showToast(`Lỗi / Error: ${j.error}`, 'danger');
          setTimeout(() => { if (banner) banner.style.display = 'none'; }, 8000);
        } else if (j.status === 'cancelled') {
          clearInterval(state.jobPollTimer);
          showToast(state.lang === 'vi' ? 'Đã hủy tác vụ' : 'Job cancelled', 'warning');
          setTimeout(() => { if (banner) banner.style.display = 'none'; }, 3000);
        }
      }
    } catch (err) {
      console.error('Job polling error:', err);
    }
  }, 1000);
}

function updateActiveJobBanner(job) {
  const title = document.getElementById('active-job-title');
  const speed = document.getElementById('active-job-speed');
  const stage = document.getElementById('active-job-stage');
  const bar = document.getElementById('active-job-bar');

  if (title) title.textContent = job.appName || job.bundleId || t('job_decrypting');
  if (speed) speed.textContent = job.speed || '';
  if (stage) stage.textContent = job.stage || '';
  if (bar) bar.style.width = `${Math.min(100, Math.max(0, job.progress || 0))}%`;
}

function updateLiveMonitor(job) {
  const nameEl = document.getElementById('log-app-name');
  const bundleEl = document.getElementById('log-app-bundle');
  const badgeEl = document.getElementById('log-app-badge');
  const bar = document.getElementById('log-progress-bar');
  const stageEl = document.getElementById('log-stage-text');
  const speedEl = document.getElementById('log-speed-text');

  if (nameEl) nameEl.textContent = job.appName || job.bundleId;
  if (bundleEl) bundleEl.textContent = `${job.bundleId} ${job.version ? `v${job.version}` : ''}`;
  if (badgeEl) badgeEl.textContent = job.status.toUpperCase();
  if (bar) bar.style.width = `${Math.min(100, Math.max(0, job.progress || 0))}%`;
  if (stageEl) stageEl.textContent = job.stage || '';
  if (speedEl) speedEl.textContent = job.speed || '';
}

async function cancelActiveJob() {
  if (!state.activeJob) return;
  try {
    await fetch('/api/decrypt/cancel', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ id: state.activeJob.id })
    });
  } catch (err) {
    console.error(err);
  }
}

// --- App Store Search ---
async function doSearch() {
  const input = document.getElementById('search-input');
  const countrySelect = document.getElementById('search-country');
  const term = input ? input.value.trim() : '';
  const country = countrySelect ? countrySelect.value : 'VN';

  if (!term) return;

  const loading = document.getElementById('search-loading');
  const empty = document.getElementById('search-empty');
  const grid = document.getElementById('search-results-grid');

  if (loading) loading.style.display = 'block';
  if (empty) empty.style.display = 'none';
  if (grid) grid.style.display = 'none';

  try {
    const res = await fetch(`/api/search?term=${encodeURIComponent(term)}&country=${country}`);
    const data = await res.json();

    if (loading) loading.style.display = 'none';

    if (data.success && data.results && data.results.length > 0) {
      renderSearchResults(data.results);
      if (grid) grid.style.display = 'grid';
    } else {
      if (empty) {
        empty.style.display = 'block';
        empty.querySelector('h3').textContent = t('search_no_results');
        empty.querySelector('p').textContent = '';
      }
    }
  } catch (err) {
    if (loading) loading.style.display = 'none';
    showToast(err.message, 'danger');
  }
}

function renderSearchResults(results) {
  const grid = document.getElementById('search-results-grid');
  if (!grid) return;

  grid.innerHTML = results.map(app => {
    const icon = app.iconUrl || '';
    const platforms = (app.platforms || ['iOS']).map(p => `<span class="badge badge-primary">${p}</span>`).join(' ');

    return `
      <div class="app-card">
        <div class="app-card-header">
          <img class="app-icon" src="${icon}" alt="${escapeHtml(app.name)}" onerror="this.src='data:image/svg+xml,<svg xmlns=%22http://www.w3.org/2000/svg%22 viewBox=%220 0 24 24%22 fill=%22none%22 stroke=%22currentColor%22><rect width=%2218%22 height=%2218%22 x=%223%22 y=%223%22 rx=%224%22/></svg>'">
          <div class="app-info">
            <h4 class="app-title" title="${escapeHtml(app.name)}">${escapeHtml(app.name)}</h4>
            <span class="app-dev">${escapeHtml(app.developer || '')}</span>
            <div class="app-meta">
              <span>v${escapeHtml(app.version || '')}</span>
              <span>•</span>
              <span class="font-mono">${escapeHtml(app.bundleId)}</span>
            </div>
            <div class="mt-2 flex-gap">${platforms}</div>
          </div>
        </div>
        <div class="app-card-actions mt-3 flex-gap">
          <button class="btn btn-primary btn-xs flex-1" onclick="selectAppForDecrypt('${escapeHtml(app.bundleId)}')">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" width="12" height="12">
              <rect x="3" y="11" width="18" height="11" rx="2" ry="2"></rect>
              <path d="M7 11V7a5 5 0 0 1 9.9-1"></path>
            </svg>
            ${t('btn_decrypt')}
          </button>
          <button class="btn btn-outline btn-xs" onclick="openVersionsFor('${escapeHtml(app.bundleId)}')">
            ${t('btn_versions')}
          </button>
        </div>
      </div>
    `;
  }).join('');
}

function selectAppForDecrypt(bundleId) {
  switchTab('decrypt');
  selectSourceMode('store');
  const targetInput = document.getElementById('decrypt-target');
  if (targetInput) targetInput.value = bundleId;
}

function openVersionsFor(bundleId) {
  switchTab('versions');
  const targetInput = document.getElementById('versions-target');
  if (targetInput) {
    targetInput.value = bundleId;
    lookupVersions();
  }
}

// --- Version History ---
async function lookupVersions() {
  const target = document.getElementById('versions-target').value.trim();
  if (!target) return;

  const loading = document.getElementById('versions-loading');
  const resultsWrap = document.getElementById('versions-results-wrap');
  const tbody = document.getElementById('versions-tbody');

  if (loading) loading.style.display = 'block';
  if (resultsWrap) resultsWrap.style.display = 'none';

  let url = `/api/versions?target=${encodeURIComponent(target)}`;
  if (state.activeAccount) url += `&accountId=${state.activeAccount.id}`;

  try {
    const res = await fetch(url);
    const data = await res.json();
    if (loading) loading.style.display = 'none';

    if (data.success && data.versions) {
      document.getElementById('versions-app-name').textContent = data.appName;
      document.getElementById('versions-app-bundle').textContent = `${data.bundleId} (ID: ${data.appId})`;
      document.getElementById('versions-count').textContent = `${data.count} ${t('versions_unit')}`;

      tbody.innerHTML = data.versions.map(v => `
        <tr>
          <td class="font-mono font-semibold">${v.versionId} ${v.current ? `<span class="badge badge-success ml-2">${t('badge_latest')}</span>` : ''}</td>
          <td>${v.displayVersion || `<button class="btn btn-outline btn-xs" onclick="fetchVersionMeta('${data.bundleId}', '${v.versionId}', this)">${t('btn_view_meta')}</button>`}</td>
          <td class="text-secondary text-xs">${v.releaseDate || '—'}</td>
          <td>
            <button class="btn btn-primary btn-xs" onclick="selectVersionForDecrypt('${data.bundleId}', '${v.versionId}')">
              ${t('btn_decrypt')}
            </button>
          </td>
        </tr>
      `).join('');

      if (resultsWrap) resultsWrap.style.display = 'block';
    } else {
      showToast(data.error || 'Failed to lookup versions', 'danger');
    }
  } catch (err) {
    if (loading) loading.style.display = 'none';
    showToast(err.message, 'danger');
  }
}

async function fetchVersionMeta(bundleId, versionId, btn) {
  btn.disabled = true;
  btn.textContent = '...';
  try {
    let url = `/api/version-metadata?target=${encodeURIComponent(bundleId)}&versionId=${versionId}`;
    if (state.activeAccount) url += `&accountId=${state.activeAccount.id}`;
    const res = await fetch(url);
    const data = await res.json();
    if (data.success) {
      const td = btn.parentElement;
      td.textContent = data.displayVersion || data.bundleVersion || 'N/A';
      const tr = td.parentElement;
      const dateTd = tr.children[2];
      if (dateTd && data.releaseDate) dateTd.textContent = data.releaseDate;
    }
  } catch (err) {
    btn.disabled = false;
  }
}

function selectVersionForDecrypt(bundleId, versionId) {
  switchTab('decrypt');
  selectSourceMode('store');
  document.getElementById('decrypt-target').value = bundleId;
  document.getElementById('decrypt-version-id').value = versionId;
}

// --- Jailbroken Device & Credential Minting ---
function toggleDeviceUSBMode() {
  const isUsb = document.getElementById('dev-usb').checked;
  const netFields = document.getElementById('dev-network-fields');
  if (netFields) {
    netFields.style.display = isUsb ? 'none' : 'block';
  }
}

async function loadDeviceConfig() {
  try {
    const res = await fetch('/api/device/config');
    const data = await res.json();
    if (data.success && data.config) {
      const c = data.config;
      document.getElementById('dev-usb').checked = !!c.usb;
      document.getElementById('dev-host').value = c.host || '';
      document.getElementById('dev-port').value = c.port || 22;
      document.getElementById('dev-user').value = c.user || 'root';
      document.getElementById('dev-password').value = c.password || '';
      document.getElementById('dev-udid').value = c.udid || '';
      toggleDeviceUSBMode();
    }
  } catch (err) {
    console.error('Failed to load device config:', err);
  }
}

async function saveDeviceConfig() {
  const payload = {
    usb: document.getElementById('dev-usb').checked,
    host: document.getElementById('dev-host').value.trim(),
    port: parseInt(document.getElementById('dev-port').value) || 22,
    user: document.getElementById('dev-user').value.trim(),
    password: document.getElementById('dev-password').value,
    udid: document.getElementById('dev-udid').value.trim()
  };

  try {
    const res = await fetch('/api/device/config', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload)
    });
    const data = await res.json();
    if (data.success) {
      showToast(state.lang === 'vi' ? 'Đã lưu cấu hình thiết bị!' : 'Device settings saved!', 'success');
    } else {
      showToast(data.error || 'Failed to save config', 'danger');
    }
  } catch (err) {
    showToast(err.message, 'danger');
  }
}

async function testDeviceConnection() {
  const resWrap = document.getElementById('device-test-results');
  const list = document.getElementById('device-test-checks-list');
  if (resWrap) resWrap.style.display = 'block';
  if (list) list.innerHTML = `<div class="p-2 text-secondary text-xs">${t('check_ssh_connecting')}</div>`;

  try {
    const res = await fetch('/api/device/test', { method: 'POST' });
    const data = await res.json();

    if (data.success && data.checks) {
      list.innerHTML = data.checks.map(c => `
        <div class="test-check-item">
          <div>
            <span class="font-semibold text-sm text-primary">${escapeHtml(c.name)}</span>
            <span class="text-xs text-secondary block">${escapeHtml(c.detail)}</span>
          </div>
          <div>
            ${c.passed ? `<span class="badge badge-success">${t('badge_ready')}</span>` : `<span class="badge badge-danger">${t('badge_missing')}</span>`}
          </div>
        </div>
      `).join('');
    } else {
      list.innerHTML = `<div class="p-2 text-danger text-xs">${escapeHtml(data.error || 'Test failed')}</div>`;
    }
  } catch (err) {
    if (list) list.innerHTML = `<div class="p-2 text-danger text-xs">${escapeHtml(err.message)}</div>`;
  }
}

function quickMintAccount(accountId) {
  switchTab('device');
  startMintCredential(accountId);
}

async function startMintCredential(specificAccountId) {
  const accountId = specificAccountId || (state.activeAccount ? state.activeAccount.id : '');
  if (!accountId) {
    showToast(state.lang === 'vi' ? 'Chưa chọn tài khoản Apple ID để mint' : 'Please select an Apple ID to mint credential', 'warning');
    return;
  }

  const modal = document.getElementById('modal-mint');
  if (modal) modal.style.display = 'flex';

  try {
    const res = await fetch('/api/device/mint', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ accountId })
    });
    const data = await res.json();

    if (modal) modal.style.display = 'none';

    if (data.success) {
      showToast(data.message || (state.lang === 'vi' ? 'Đã tạo chứng chỉ FairPlay thành công!' : 'FairPlay credential minted successfully!'), 'success');
      loadAccounts();
    } else {
      showToast(data.error || 'Mint failed', 'danger');
    }
  } catch (err) {
    if (modal) modal.style.display = 'none';
    showToast(err.message, 'danger');
  }
}

// --- Decrypted Files Browser ---
async function loadDecryptedFiles() {
  try {
    const res = await fetch('/api/files');
    const data = await res.json();
    if (data.success) {
      renderFilesTable(data.files || []);
      const badge = document.getElementById('files-badge');
      if (badge) {
        badge.textContent = (data.files || []).length;
        badge.style.display = data.files && data.files.length > 0 ? 'inline-block' : 'none';
      }
    }
  } catch (err) {
    console.error('Failed to load files:', err);
  }
}

function renderFilesTable(files) {
  const tbody = document.getElementById('files-tbody');
  const empty = document.getElementById('files-empty');
  const table = document.getElementById('files-table');

  if (!tbody) return;

  if (files.length === 0) {
    if (empty) empty.style.display = 'block';
    if (table) table.style.display = 'none';
    return;
  }

  if (empty) empty.style.display = 'none';
  if (table) table.style.display = 'table';

  tbody.innerHTML = files.map(f => `
    <tr>
      <td>
        <div class="font-semibold text-primary text-sm">${escapeHtml(f.name)}</div>
        ${f.bundleId ? `<span class="text-xs text-secondary font-mono">${escapeHtml(f.bundleId)} (v${escapeHtml(f.version || '')})</span>` : ''}
      </td>
      <td class="text-sm font-mono">${f.sizeHuman}</td>
      <td class="text-sm text-secondary">${f.modifiedDate}</td>
      <td>
        <div class="flex-gap">
          <a class="btn btn-outline btn-xs" href="/api/files/download?name=${encodeURIComponent(f.name)}" download>
            ${t('btn_download')}
          </a>
          <button class="btn btn-danger btn-xs" onclick="deleteFile('${escapeHtml(f.name)}')">
            ${t('btn_delete')}
          </button>
        </div>
      </td>
    </tr>
  `).join('');
}

async function openOutputFolder() {
  try {
    await fetch('/api/files/open', { method: 'POST' });
  } catch (err) {
    showToast(err.message, 'danger');
  }
}

async function deleteFile(name) {
  const confirmMsg = state.lang === 'vi' ? `Bạn có chắc muốn xóa file "${name}"?` : `Are you sure you want to delete "${name}"?`;
  if (!confirm(confirmMsg)) return;

  try {
    const res = await fetch('/api/files/delete', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ name })
    });
    const data = await res.json();
    if (data.success) {
      showToast(state.lang === 'vi' ? 'Đã xóa file' : 'File deleted', 'success');
      loadDecryptedFiles();
    } else {
      showToast(data.error || 'Failed to delete file', 'danger');
    }
  } catch (err) {
    showToast(err.message, 'danger');
  }
}

// --- Utilities ---
function escapeHtml(str) {
  if (!str) return '';
  return String(str)
    .replace(/&/g, '&amp;')
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;')
    .replace(/"/g, '&quot;')
    .replace(/'/g, '&#039;');
}

function formatBytes(bytes) {
  if (bytes === 0 || !bytes) return '0 B';
  const k = 1024;
  const sizes = ['B', 'KB', 'MB', 'GB'];
  const i = Math.floor(Math.log(bytes) / Math.log(k));
  return parseFloat((bytes / Math.pow(k, i)).toFixed(1)) + ' ' + sizes[i];
}
