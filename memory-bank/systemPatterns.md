# System Patterns: Go Cronjob Package

## 1. Tổng quan kiến trúc

Hệ thống được thiết kế theo các nguyên tắc **Clean Architecture** và **SOLID**, nhằm đảm bảo tính dễ bảo trì, dễ mở rộng và khả năng tái sử dụng cao. Mục tiêu là tạo ra một package mà người dùng có thể dễ dàng tích hợp và sử dụng, đồng thời cũng dễ dàng cho các nhà phát triển đóng góp và mở rộng.

Kiến trúc ban đầu sẽ tập trung vào giải pháp **memory-first**, nghĩa là tất cả thông tin về job và lịch trình của chúng sẽ được lưu trữ và quản lý trong bộ nhớ của ứng dụng khi nó đang chạy.

## 2. Các thành phần chính (Components)

Kiến trúc có thể được chia thành các thành phần chính sau:

1.  **Job Definition Handler (Bộ xử lý định nghĩa Job):**
    *   **Trách nhiệm:** Tiếp nhận yêu cầu đăng ký job từ người dùng thông qua các hàm/phương thức của API Layer. Xác thực thông tin đầu vào (cron expression, hàm thực thi) và chuẩn bị dữ liệu để đưa vào `Job Registry`.
    *   **Pattern có thể áp dụng:** Builder Pattern (để tạo đối tượng job một cách linh hoạt), Validation (để kiểm tra tính hợp lệ của đầu vào).

2.  **Job Registry (Bộ đăng ký Job):**
    *   **Trách nhiệm:** Lưu trữ danh sách các job đã được phát hiện và phân tích. Mỗi entry trong registry sẽ chứa thông tin về hàm cần thực thi và lịch trình của nó.
    *   **Pattern có thể áp dụng:** Singleton Pattern (để đảm bảo chỉ có một registry duy nhất trong ứng dụng), Repository Pattern (để trừu tượng hóa việc lưu trữ và truy xuất job, ban đầu là in-memory).

3.  **Scheduler (Bộ lập lịch):**
    *   **Trách nhiệm:** Dựa trên cron expression của mỗi job trong registry, quyết định thời điểm thực thi job. Quản lý vòng đời của các goroutine thực thi job.
    *   **Pattern có thể áp dụng:** Observer Pattern (để thông báo cho các thành phần khác khi job được thực thi hoặc thay đổi trạng thái), Worker Pool Pattern (để quản lý số lượng goroutine thực thi job đồng thời).

4.  **Job Executor (Bộ thực thi Job):**
    *   **Trách nhiệm:** Thực thi hàm (function) tương ứng với job khi đến lịch. Xử lý lỗi và logging trong quá trình thực thi.
    *   **Pattern có thể áp dụng:** Command Pattern (để đóng gói yêu cầu thực thi job thành một đối tượng).

5.  **API Layer (Lớp giao diện lập trình):**
    *   **Trách nhiệm:** Cung cấp các hàm public để người dùng tương tác với hệ thống cron job (ví dụ: khởi tạo, dừng, thêm job thủ công nếu cần).
    *   **Pattern có thể áp dụng:** Facade Pattern (để cung cấp một giao diện đơn giản cho một hệ thống con phức tạp).

## 3. Luồng hoạt động chính (Main Flow)

1.  **Khởi tạo:** Người dùng gọi hàm khởi tạo của package.
2.  **Tiếp nhận Định nghĩa:** Người dùng gọi hàm/phương thức trong `API Layer` để định nghĩa một job, cung cấp cron expression và hàm cần thực thi.
3.  **Xử lý và Đăng ký Job:** `Job Definition Handler` tiếp nhận thông tin, xác thực và sau đó `Job Registry` lưu trữ thông tin job.
4.  **Lập lịch:** `Scheduler` đọc thông tin từ `Job Registry`, tính toán thời điểm thực thi cho mỗi job và chuẩn bị các goroutine.
5.  **Thực thi:** Khi đến thời điểm, `Scheduler` kích hoạt `Job Executor` để chạy hàm tương ứng của job trong một goroutine mới.
6.  **Logging và Xử lý lỗi:** Quá trình thực thi được log lại, và các lỗi (nếu có) được xử lý theo cơ chế đã định.

## 4. Quyết định kỹ thuật quan trọng

-   **Memory-First:** Giải pháp ban đầu sẽ hoàn toàn dựa vào bộ nhớ. Điều này giúp đơn giản hóa việc triển khai và phù hợp với các ứng dụng không yêu cầu tính bền vững (persistence) của job qua các lần khởi động lại ứng dụng hoặc khi ứng dụng bị crash.
-   **Sử dụng Goroutines:** Tận dụng tối đa khả năng xử lý đồng thời của Golang thông qua goroutines để thực thi các job một cách bất đồng bộ và hiệu quả.
-   **Cron Expression Standard:** Sử dụng thư viện chuẩn hoặc một thư viện phổ biến, được kiểm thử kỹ lưỡng để phân tích và xử lý cron expression, đảm bảo tính chính xác và linh hoạt trong việc định nghĩa lịch trình.
-   **Cấu hình tối thiểu:** Hướng tới việc người dùng chỉ cần một dòng lệnh để định nghĩa job mà không cần thêm bất kỳ cấu hình phức tạp nào khác.

## 5. Khả năng mở rộng

Kiến trúc được thiết kế để có thể mở rộng trong tương lai:

-   **Hỗ trợ Persistent Storage:** `Job Registry` có thể được mở rộng để sử dụng các giải pháp lưu trữ bền vững như database (SQL, NoSQL) hoặc file system, thay vì chỉ in-memory.
-   **Tích hợp Message Queues:** `Scheduler` và `Job Executor` có thể được điều chỉnh để làm việc với các hệ thống message queue (ví dụ: RabbitMQ, Kafka, Redis Streams). Job có thể được đẩy vào queue và được xử lý bởi các worker riêng biệt, tăng khả năng chịu lỗi và phân tán tải.
-   **Giao diện quản lý (UI):** Có thể phát triển một giao diện web hoặc CLI để theo dõi và quản lý các job.
-   **Plugin System:** Cho phép người dùng mở rộng chức năng thông qua các plugin (ví dụ: custom logger, custom error handler).
-   **Tích hợp Giám sát (Monitoring Integration):** Hệ thống có thể được mở rộng để expose metrics (ví dụ: số lượng job đang chạy, số job lỗi, thời gian thực thi job) theo một định dạng chuẩn (ví dụ: Prometheus exposition format) để các công cụ giám sát như Grafana có thể thu thập và hiển thị.