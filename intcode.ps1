
function Invoke-IntCode {

    [CMDletBinding()]
    Param (
        [Int[]]
        $Intcode
    )

    $instructionPointer = 0
    while ($true) {
        $opcode = $Intcode[$instructionPointer]
        Write-Verbose -Message "instructionPointer: $instructionPointer"
        Write-Verbose -Message "Opcode: $opcode"

        switch ($opcode) {

            # Addition
            1 {
                $instructionSize = 4

                $p1 = $Intcode[$instructionPointer + 1]
                $p2 = $Intcode[$instructionPointer + 2]
                $p3 = $Intcode[$instructionPointer + 3]

                $Intcode[$p3] = $Intcode[$p1] + $Intcode[$p2]
            }

            2 {
                $instructionSize = 4

                $p1 = $Intcode[$instructionPointer + 1]
                $p2 = $Intcode[$instructionPointer + 2]
                $p3 = $Intcode[$instructionPointer + 3]

                $Intcode[$p3] = $Intcode[$p1] * $Intcode[$p2]
            }

            99 {
                $instructionSize = 1
                return $Intcode
            }

            default {
                throw "invalid opcode: $opcode"
            }
        }

        $instructionPointer += $instructionSize
    }
}
