cd D:\dev\Go\src\AlgebraCalculator\ruleFiles
fyne bundle simpRulesExpand.txt > bundled.go
fyne bundle -append simpRulesSumUp.txt >> bundled.go
move "bundled.go" "D:\dev\Go\src\AlgebraCalculator\app\fyne\AlgebraCalculator"