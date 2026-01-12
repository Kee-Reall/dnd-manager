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
    up_file = f"{base_name}.up.sql"
    down_file = f"{base_name}.down.sql"
    files = [up_file, down_file]

    for file in files:
        path=os.path.join(scripts_dir, file)
        with open(path, 'w') as f:
            f.write("")

    print(f"{up_file}\n{down_file}\nhas been created in: {scripts_dir}/")

if __name__ == "__main__":
    main()