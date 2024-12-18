import os
import shutil


def restructure_code(base_dir, output_dir):
    for root, dirs, files in os.walk(base_dir):
        for file in files:
            # Kiểm tra file có định dạng lcXXX
            if file.startswith("lc"):
                # Lấy ngày từ đường dẫn gốc
                relative_path = os.path.relpath(root, base_dir)
                date = relative_path.split(os.sep)[0]

                # Tách năm, tháng, ngày từ tên thư mục
                try:
                    year, month, day = date.split("-")
                except ValueError:
                    print(f"Bỏ qua thư mục không đúng định dạng ngày: {relative_path}")
                    continue

                # Tạo thư mục mới theo định dạng /docs/code/{{năm}}-{{tháng}}/ngày
                new_dir = os.path.join(output_dir, f"{year}-{month}", day)
                os.makedirs(new_dir, exist_ok=True)

                # Di chuyển tệp vào thư mục mới
                src_file = os.path.join(root, file)
                dst_file = os.path.join(new_dir, file)
                shutil.move(src_file, dst_file)


# Cấu hình thư mục gốc và thư mục đích
base_directory = "./code"  # Đường dẫn thư mục cũ
output_directory = "./code"  # Đường dẫn thư mục mới

# Gọi hàm thực hiện tái cấu trúc
restructure_code(base_directory, output_directory)

print("Đã tái cấu trúc xong thư mục và tệp tin.")
