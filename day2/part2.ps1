
. "$PSScriptRoot/../intcode.ps1"

foreach ($noun in 0..99) {
    "Noun: $noun"
    foreach ($verb in 0..99) {
        $intcode = (Get-Content "$PSScriptRoot/input.txt") -split ','

        $intcode[1] = $noun
        $intcode[2] = $verb

        $result = Invoke-IntCode $intcode

        if ($result[0] -eq 19690720) {
            "Verb: $verb"
            100 * $noun + $verb
            return
        }
    }
}

