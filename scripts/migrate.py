"""
Скрипт для создания файлов миграций.
Использование: python migrategen.py <имя_миграции>
"""
import sys
import os
from datetime import datetime

def main():
    if len(sys.argv) != 2:
        print("Использование: python migrategen.py <имя_миграции>\nПример: python migrategen.py add_index_for_user_table")
        sys.exit(1)

    migration_name = sys.argv[1]

    scripts_dir = "migrations"
    if not os.path.exists(scripts_dir):
        os.makedirs(scripts_dir)

    timestamp = datetime.now().strftime("%Y%m%d%H%M%S")
    base_name = f"{timestamp}_{migration_name}"
    files = [f"{base_name}.up.sql", f"{base_name}.down.sql"]

    for file in files:
        with open(os.path.join(scripts_dir, file), 'w') as f:
            f.write("")

    print(f"{files[0]}\n{files[1]}\nhas been created in: {scripts_dir}/")

if __name__ == "__main__":
    main()