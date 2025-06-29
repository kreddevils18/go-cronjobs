# Project Brief: Go Cronjob Package

## 1. Giới thiệu

Dự án này nhằm mục đích xây dựng một package Golang cho phép người dùng định nghĩa và quản lý các công việc (job) xử lý bất đồng bộ một cách nhanh chóng và dễ dàng. Điểm đặc biệt là các job này có thể được cấu hình và khởi tạo chỉ thông qua một dòng lệnh (command) duy nhất trong code, giúp giảm thiểu boilerplate code và tăng tốc độ phát triển.

Package này được thiết kế để có thể tái sử dụng ở nhiều dự án và repository khác nhau.

## 2. Mục tiêu chính

- **Phát triển nhanh chóng:** Cho phép định nghĩa và chạy các cron job với cấu hình tối thiểu.
- **Dễ sử dụng:** Giao diện lập trình (API) đơn giản, trực quan.
- **Tái sử dụng cao:** Dễ dàng tích hợp vào các dự án Golang khác nhau.
- **Mở rộng linh hoạt:** Kiến trúc cho phép mở rộng các tính năng trong tương lai (ví dụ: hỗ trợ các loại queue khác nhau, persistent storage).

## 3. Yêu cầu cốt lõi

- **Định nghĩa job qua dòng lệnh:** Hệ thống phải cung cấp một cách thức đơn giản, ví dụ như một hàm hoặc phương thức, để người dùng có thể định nghĩa và đăng ký cron job chỉ bằng một dòng lệnh trong code.
- **Xử lý bất đồng bộ:** Các job phải được thực thi bất đồng bộ, không làm ảnh hưởng đến luồng chính của ứng dụng.
- **Quản lý vòng đời job:** Cung cấp cơ chế để khởi động, dừng và theo dõi trạng thái của các job.
- **Clean Code:** Mã nguồn phải được tổ chức rõ ràng, dễ hiểu, dễ bảo trì và tuân thủ các nguyên tắc clean code.
- **Memory-first:** Phiên bản đầu tiên sẽ lưu trữ và quản lý job trong bộ nhớ (in-memory).

## 4. Phạm vi dự án (Ban đầu)
- Phát triển core engine để xử lý các định nghĩa job qua dòng lệnh và quản lý job.
- Implement cơ chế chạy job bất đồng bộ cơ bản.
- Cung cấp API cơ bản để tương tác với hệ thống cron job.
- Tập trung vào giải pháp memory-first.
- **Tích hợp giám sát (Monitoring):** Khả năng tích hợp với các công cụ giám sát như Grafana để theo dõi hiệu suất và trạng thái của các job.

## 5. Tiêu chí thành công

- Người dùng có thể tích hợp và sử dụng package để chạy cron job với nỗ lực tối thiểu.
- Package hoạt động ổn định và đáng tin cậy.
- Kiến trúc dễ dàng cho việc mở rộng và bảo trì sau này.