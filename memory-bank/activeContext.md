# Active Context: Go Cronjob Package

## 1. Trạng thái hiện tại

-   **Giai đoạn:** Khởi tạo dự án và thiết lập nền tảng ban đầu.
-   **Công việc vừa hoàn thành:**
    -   Tạo các tài liệu cơ bản trong Memory Bank:
        -   `projectbrief.md`: Mô tả tổng quan dự án.
        -   `productContext.md`: Bối cảnh sản phẩm, vấn đề giải quyết.
        -   `systemPatterns.md`: Kiến trúc hệ thống và các mẫu thiết kế dự kiến.
        -   `techContext.md`: Bối cảnh kỹ thuật, công nghệ sử dụng.
-   **Tập trung hiện tại:** Hoàn thiện bộ tài liệu Memory Bank ban đầu.

## 2. Thay đổi gần đây

-   **Thay đổi phương thức định nghĩa job:** Chuyển từ việc sử dụng comment đặc biệt sang sử dụng một dòng lệnh (command) duy nhất trong code để định nghĩa job. Điều này ảnh hưởng đến `Parser` (đã được thay thế/điều chỉnh thành một `Job Definition Handler` nhận input từ API) và cách người dùng tương tác với package.
-   **Thư viện Logging:** Quyết định sử dụng `github.com/kreddevils18/go-logger` thay vì package `log` tiêu chuẩn của Go.
-   **Thư viện Cấu hình:** Quyết định sử dụng `github.com/spf13/viper` để quản lý cấu hình, với định dạng ưu tiên là YML.
-   Cập nhật `projectbrief.md` để phản ánh thay đổi về phương thức định nghĩa job. (Hoàn thành)
-   Cập nhật `productContext.md` để phản ánh thay đổi về phương thức định nghĩa job và lợi ích. (Hoàn thành)
-   Cập nhật `systemPatterns.md` để thay thế/điều chỉnh `Parser` thành `Job Definition Handler` và cập nhật luồng hoạt động. (Hoàn thành)
-   Cập nhật `techContext.md` để phản ánh thay đổi về `Reflection` (không còn cần thiết), `Parser` (thay đổi vai trò), ghi nhận việc sử dụng `go-logger` và `go-viper`. (Hoàn thành)

## 3. Các bước tiếp theo (Next Steps)

1.  **Tạo `progress.md`:**
    -   Mô tả những gì đã hoạt động (hiện tại là chưa có gì vì mới bắt đầu).
    -   Liệt kê những gì cần xây dựng.
    -   Xác định trạng thái hiện tại của dự án (ví dụ: "Giai đoạn lên kế hoạch và thiết kế").
    -   Ghi nhận các vấn đề đã biết (nếu có).
2.  **Tạo file `.cursorrules` (nếu cần):** Ghi lại các quy tắc hoặc hướng dẫn cụ thể cho AI (Cursor) trong quá trình phát triển dự án này.
3.  **Bắt đầu phát triển các thành phần cốt lõi của package:**
    -   Thiết kế và implement `Job Definition Handler`.
    -   Xây dựng `Job Registry` (in-memory).
    -   Implement `Scheduler` cơ bản.
    -   Tạo `Job Executor`.
    -   Xây dựng API Layer ban đầu (bao gồm hàm để đăng ký job).
    -   Tích hợp `go-viper` để đọc cấu hình (ví dụ: cấu hình logging, tùy chọn scheduler).
4.  **Viết Unit Tests:** Đảm bảo các thành phần cốt lõi được kiểm thử kỹ lưỡng.
5.  **Tạo ví dụ sử dụng (Example Usage):** Xây dựng một ứng dụng nhỏ để minh họa cách sử dụng package.

## 4. Quyết định và cân nhắc đang hoạt động

-   **Ưu tiên Memory-First:** Tập trung hoàn thiện giải pháp chạy cron job trong bộ nhớ trước khi xem xét các giải pháp lưu trữ bền vững hoặc tích hợp queue.
-   **Clean Code và Test Coverage:** Đặt ưu tiên cao cho việc viết mã sạch, dễ hiểu và có độ bao phủ test tốt ngay từ đầu.
-   **API Định nghĩa Job:** Cần sớm định nghĩa API để đăng ký job một cách rõ ràng và đơn giản. Ví dụ:
    ```go
    cronjob.Register("*/5 * * * *", myAwesomeJob, "MyAwesomeJob")
    ```
-   **Xử lý lỗi và Logging:** Cần thiết kế cơ chế logging và xử lý lỗi linh hoạt, cho phép người dùng dễ dàng theo dõi và gỡ lỗi.
-   **Giám sát (Monitoring):** Nghiên cứu và cân nhắc việc tích hợp Prometheus để thu thập metrics và Grafana để trực quan hóa, giúp theo dõi hiệu suất và trạng thái của các job.

## 5. Rủi ro tiềm ẩn

-   **Thiết kế API:** API cần được thiết kế cẩn thận để dễ sử dụng và mở rộng trong tương lai.
-   **Quản lý Goroutine:** Nếu không cẩn thận, việc tạo và quản lý goroutine cho mỗi job có thể dẫn đến resource leak hoặc contention.
-   **Phạm vi ban đầu:** Cần giữ phạm vi của phiên bản đầu tiên ở mức quản lý được, tránh ôm đồm quá nhiều tính năng.