
$input = Get-Content "$PSScriptRoot/input.txt"

# $input = @(12, 14, 1969, 100756)
# $input = 100756 #1969

function Get-RequiredFuel ($i) {
    [Math]::Floor($i / 3) - 2
}

$total = 0

foreach ($i in $input) {
    $fuel = Get-RequiredFuel -i $i
    while ($fuel -gt 0) {
        $total += $fuel
        $fuel = Get-RequiredFuel -i $fuel
    }
}

$total
