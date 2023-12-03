
for ($i=2; $i -le 2; $i++) {
    $dirName = "day$i"
    $destinationTest = Join-Path $dirName "main_test.go"
    $destinationMain = Join-Path $dirName "main.go"

    New-Item -ItemType Directory -Name $dirName
    Copy-Item -Path "templates/main_test.go" -Destination $destinationTest
    Copy-Item -Path "templates/main.go" -Destination $destinationMain
    New-Item -ItemType File -Name "input.txt" -Path "$dirName"
    New-Item -ItemType File -Name "input_example_part1.txt" -Path "$dirName"
    New-Item -ItemType File -Name "input_example_part2.txt" -Path "$dirName"


    $content = Get-Content $destinationTest
    $content[0] = "package "+$dirName
    $content[16] = "func TestDay$i(t *testing.T) {"
    $content | Set-Content $destinationTest

    $content = Get-Content $destinationMain
    $content[0] = "package "+$dirName
    $content | Set-Content $destinationMain

}
