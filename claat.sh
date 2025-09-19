#!/bin/bash

# Array of Markdown files to process with claat
markdown_files=(
    "./generics/index.md"
    "./delve/delve.md"
    "./suggestedfix/index.md"
    "./assembly/index.md"
    "./toolexec/index.md"
)

# Loop through each Markdown file and process with claat
for md_file in "${markdown_files[@]}"; do
    echo "Processing: $md_file"

    # Check if file exists
    if [ -f "$md_file" ]; then
        go tool claat export -o docs "$md_file"
        echo "✓ Exported: $md_file"
    else
        echo "✗ File not found: $md_file"
    fi

    echo "---"
done

echo "All files processed!"
