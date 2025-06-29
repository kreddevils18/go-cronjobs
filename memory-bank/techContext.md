# Tech Context: Go Cronjob Package

## 1. Ngôn ngữ và Môi trường

-   **Ngôn ngữ chính:** Go (Golang).
    -   Lý do: Hiệu năng cao, hỗ trợ tốt cho lập trình đồng thời (concurrency) với goroutines và channels, hệ sinh thái thư viện phong phú, biên dịch ra mã máy (native binary) giúp dễ dàng triển khai.
-   **Phiên bản Go:** Sử dụng phiên bản Go ổn định mới nhất (ví dụ: Go 1.18+ để tận dụng generics nếu cần, hoặc phiên bản LTS phù hợp).
-   **Hệ điều hành:** Package được thiết kế để tương thích với các hệ điều hành phổ biến hỗ trợ Go (Linux, macOS, Windows).
-   **Quản lý Dependencies:** Go Modules.

## 2. Thư viện chính (Core Libraries)

-   **Cron Expression Parsing & Scheduling:**
    -   Xem xét sử dụng một thư viện cron uy tín và được cộng đồng sử dụng rộng rãi như `robfig/cron` (ví dụ: `github.com/robfig/cron/v3`). Thư viện này cung cấp các chức năng mạnh mẽ để phân tích cú pháp cron expression và quản lý lịch trình.
    -   Lựa chọn này giúp giảm thiểu việc phải tự xây dựng lại logic phức tạp của việc xử lý cron, đảm bảo tính chính xác và độ tin cậy.
-   **Function Handling:**
    -   Người dùng sẽ truyền trực tiếp tham chiếu đến hàm (function reference) khi định nghĩa job. Điều này giúp loại bỏ sự cần thiết của `reflect` để tìm hàm, tăng tính tường minh và an toàn kiểu (type safety).
-   **Logging:**
    -   Sẽ sử dụng thư viện `github.com/kreddevils18/go-logger` cho việc ghi log. <mcreference link="https://github.com/kreddevils18/go-logger" index="0">0</mcreference>
    -   Thư viện này được xây dựng trên `zap` của Uber, cung cấp logging có cấu trúc, nhiều cấp độ log, và có thể cấu hình output (console, JSON, file). <mcreference link="https://github.com/kreddevils18/go-logger" index="0">0</mcreference>
    -   Điều này cho phép logging linh hoạt và mạnh mẽ hơn so với package `log` tiêu chuẩn. <mcreference link="https://github.com/kreddevils18/go-logger" index="0">0</mcreference>
-   **Configuration Management (Quản lý cấu hình):**
    -   Sử dụng thư viện `github.com/spf13/viper` (go-viper) để quản lý cấu hình.
    -   Viper hỗ trợ đọc cấu hình từ nhiều nguồn khác nhau (JSON, TOML, YAML, HCL, INI files, environment variables, remote K/V stores) và cung cấp cơ chế ghi đè cấu hình linh hoạt.
    -   Trong dự án này, sẽ tập trung vào việc sử dụng file YML cho cấu hình.

## 3. Công cụ phát triển (Development Tools)

-   **IDE:** Bất kỳ IDE nào hỗ trợ tốt cho Go (ví dụ: GoLand, VS Code với Go extension).
-   **Version Control:** Git, GitHub/GitLab/Bitbucket.
-   **Testing:**
    -   Sử dụng package `testing` tiêu chuẩn của Go để viết unit test và integration test.
    -   Cân nhắc sử dụng các thư viện hỗ trợ testing như `testify/assert` và `testify/mock` để viết test dễ dàng hơn.
-   **Linting & Formatting:**
    -   `gofmt` và `goimports` để đảm bảo code được định dạng nhất quán.
    -   `golangci-lint` để phân tích tĩnh mã nguồn, phát hiện các vấn đề tiềm ẩn và đảm bảo chất lượng code.
-   **Build & CI/CD:**
    -   Sử dụng `go build` để biên dịch package.
    -   Thiết lập quy trình CI/CD (ví dụ: GitHub Actions, GitLab CI) để tự động chạy test, linting và build khi có thay đổi trong mã nguồn.

## 4. Thiết lập môi trường phát triển

-   Cài đặt phiên bản Go phù hợp.
-   Cấu hình Go Modules cho dự án.
-   Cài đặt các công cụ phát triển cần thiết (IDE, linters).
-   Clone repository dự án.

## 5. Ràng buộc kỹ thuật (Technical Constraints)

-   **Memory-First (Ban đầu):** Phiên bản đầu tiên sẽ không có cơ chế lưu trữ bền vững (persistent storage) cho các job. Nếu ứng dụng khởi động lại, các job sẽ cần được đăng ký lại thông qua việc thực thi lại các dòng lệnh định nghĩa job.
-   **API Design:** API để định nghĩa job cần được thiết kế đơn giản, trực quan và dễ sử dụng, ví dụ: `cronjob.Register(schedule string, taskFunc func(), jobName string, params ...interface{})`.
-   **Quản lý Goroutine:** Cần quản lý cẩn thận vòng đời của các goroutine thực thi job để tránh rò rỉ tài nguyên (goroutine leaks).
-   **Xử lý lỗi:** Cung cấp cơ chế xử lý lỗi rõ ràng và cho phép người dùng tùy chỉnh cách xử lý lỗi (ví dụ: retry, log, thông báo).

## 6. Dependencies bên ngoài (External Dependencies)

-   **Tối thiểu hóa dependencies:** Cố gắng giữ số lượng thư viện bên ngoài ở mức tối thiểu để giảm độ phức tạp và tăng tính ổn định của package.
-   **Lựa chọn thư viện uy tín:** Chỉ sử dụng các thư viện được cộng đồng tin dùng, có tài liệu tốt và được bảo trì tích cực.

## 7. Giám sát và Quan sát (Monitoring & Observability)

-   **Thu thập Metrics:** Cân nhắc sử dụng thư viện như `prometheus/client_golang` để định nghĩa và expose các metrics của ứng dụng (ví dụ: số lượng job đã thực thi, số job lỗi, thời gian xử lý trung bình, số lượng goroutine đang hoạt động cho jobs).
-   **Hiển thị Metrics:** Các metrics được expose có thể được thu thập bởi một Prometheus server và sau đó được trực quan hóa bằng Grafana, cho phép người dùng theo dõi tình trạng hoạt động và hiệu năng của hệ thống cron job.

## 8. Cân nhắc về hiệu năng

-   Việc đăng ký job qua lệnh gọi hàm là trực tiếp và xảy ra khi mã đó được thực thi, thường là lúc khởi tạo ứng dụng hoặc khi logic nghiệp vụ yêu cầu.
-   Việc thực thi job trong goroutine cần được tối ưu để không tạo ra quá nhiều goroutine không cần thiết, đặc biệt với các job chạy thường xuyên.