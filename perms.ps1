$files=(git ls-files .\**\wait-for)+(git ls-files .\**\*.sh)
foreach ($file in $files) {
    git update-index --chmod=+x $file
}
