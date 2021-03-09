cd ../assets
fyne bundle simpRulesExpand.txt > bundled.go
fyne bundle -append simpRulesSumUp.txt >> bundled.go
move "bundled.go" "../AlgebraCalculator/"