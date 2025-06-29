# Progress: Go Cronjob Package

## 1. Những gì đã hoạt động (What Works)

-   Hiện tại, chưa có thành phần nào của package được implement và hoạt động. Dự án đang ở giai đoạn khởi tạo và lập kế hoạch.
-   Hoàn thiện cập nhật tài liệu Memory Bank để phản ánh các thay đổi lớn trong thiết kế (phương thức định nghĩa job, thư viện logging và cấu hình).
    -   `projectbrief.md`
    -   `productContext.md`
    -   `systemPatterns.md`
    -   `techContext.md`
    -   `activeContext.md`
    -   `progress.md` (tài liệu này)

## 2. Những gì cần xây dựng (What's Left to Build)

Đây là danh sách các hạng mục chính cần được phát triển cho phiên bản đầu tiên (memory-first):

-   **Core Engine:**
-   **Job Definition Handler:** Thiết kế và implement.
-   **Configuration Management:** Tích hợp `go-viper` để đọc cấu hình từ file YML.
        -   [ ] Định nghĩa cú pháp comment cho cron job.
        -   [ ] Implement logic quét file/package để tìm comment.
-   **Job Definition Handler:** Xử lý việc đăng ký job thông qua API.
    -   [ ] **Job Registry:**
        -   [ ] Thiết kế cấu trúc dữ liệu để lưu trữ thông tin job (in-memory).
        -   [ ] Implement các hàm để thêm, xóa, lấy thông tin job.
    -   [ ] **Scheduler:**
        -   [ ] Tích hợp thư viện cron (ví dụ: `robfig/cron`).
        -   [ ] Implement logic đăng ký job từ `Job Registry` vào scheduler.
        -   [ ] Quản lý vòng đời của các scheduled job.
    -   [ ] **Job Executor:**
        -   [ ] Implement logic thực thi hàm (function) của job trong goroutine.
        -   [ ] Cơ chế xử lý panic/error cơ bản trong quá trình thực thi job.
        -   [ ] Logging cơ bản cho việc bắt đầu và kết thúc job.
-   **API Layer:**
    -   [ ] Hàm khởi tạo chính của package (ví dụ: `cronjob.Start()`).
    -   [ ] Hàm dừng tất cả các job (ví dụ: `cronjob.Stop()`).
    -   [ ] (Tùy chọn) Các hàm để thêm/xóa job một cách thủ công (programmatically).
-   **Testing:**
    -   [ ] Unit tests cho `Parser`.
    -   [ ] Unit tests cho `Job Registry`.
    -   [ ] Unit tests cho `Scheduler` (khó hơn, có thể cần mock time).
    -   [ ] Unit tests cho `Job Executor`.
    -   [ ] Integration tests cho luồng hoạt động từ đầu đến cuối.
-   **Documentation & Examples:**
    -   [ ] Cập nhật `README.md` với hướng dẫn sử dụng chi tiết.
    -   [ ] Tạo một hoặc nhiều ví dụ sử dụng đơn giản để minh họa cách tích hợp và sử dụng package.
-   **Tooling & CI/CD:**
    -   [ ] Thiết lập `golangci-lint`.
    -   [ ] Thiết lập quy trình CI cơ bản (ví dụ: GitHub Actions) để chạy tests và linting.
-   **Monitoring & Observability (Cân nhắc cho tương lai/Phiên bản sau):**
    -   [ ] Nghiên cứu tích hợp Prometheus để thu thập metrics.
    -   [ ] Thiết lập dashboard Grafana cơ bản để hiển thị metrics.

## 3. Trạng thái hiện tại (Current Status)

-   **Giai đoạn:** Lên kế hoạch và Thiết kế ban đầu (Planning and Initial Design).
-   Bắt đầu implement `Job Definition Handler` và `Job Registry`.
-   Nghiên cứu cách tích hợp `go-viper` cho việc đọc cấu hình của package.

## 4. Vấn đề đã biết (Known Issues)

-   Chưa có vấn đề nào được xác định ở giai đoạn này vì quá trình implement chưa bắt đầu.