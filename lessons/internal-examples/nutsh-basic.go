package main

import (
	. "morr.cc/nutsh.git/dsl"
	"morr.cc/nutsh.git/dsl/bash"
)

func main() {
	Spawn("bash")

	gotoJail()
	bash.Execute("mkdir schuhkarton")
	bash.Execute("mkdir schrank")
	bash.Execute("touch schrank/jacke")
	bash.Execute("touch schrank/hut")
	bash.Execute("touch linker_schuh")

	Say("Hallo! Willkommen in der Nut-Shell! Ich möchte dir zeigen, wie du mithilfe der Kommandozeile schnell und einfach mit Dateien und Verzeichnissen umgehen kannst.")
	Say("Legen wir gleich los: Tipp mal `ls` ein und drück Enter.")
	for Prompt() {
		Output()
		if Command("^1s\n$") {
			Say("Das ist ein kleines L, keine Eins! Probier's nochmal!")
		}
		if Command("^ls\n$") {
			Say("Genau. `ls` steht kurz für \"list\" und zeigt dir die Dateien und Verzeichnisse an, die sich in deinem \"aktuellen\" Verzeichnis befinden. Die Verzeichnisse werden dabei blau dargestellt.")
			Say("Du bist gerade in einem Ordner namens `/tmp/nutsh` - das steht auch in dem blauen Text, den wir \"Prompt\" nennen. Der Prompt endet mit einem Dollarzeichen, das heißt soviel wie: \"Du kannst jetzt ein Kommando eingeben!\"")
			break
		}
	}

	Say("Du hast vielleicht gesehen, dass sich hier ein Verzeichnis namens `schrank` befindet. Um dieses zu deinem aktuellen Verzeichnis zu machen, tippst du `cd`, dann ein Leerzeichen und dann den Namen des Verzeichnisses, in das du möchtest. Begib dich doch mal \"in den Schrank\" und sieh dich darin um!")
	for Prompt() {
		Output()
		if bash.Test("$(pwd) = \"$ROOT/schrank\"") {
			if Command("^ls\n$") {
				Say("Genau. Hast du bemerkt, wie sich der Prompt geändert hat?")
				break
			}
		}
	}

	Say("Und wenn du wieder aus dem Schrank herausmöchtest? Die Abkürzung für das Verzeichnis oberhalb des aktuellen ist \"`..`\"!")
	for Prompt() {
		Output()
		if Command("^\\.\\.\n$") {
			Say("`..` ist der Name des Verzeichnisses, du musst noch dazusagen, was du damit machen möchtest. Um \"hinzugehen\", schreib `cd` davor.")
		}
		if Command("^cd\\.\\.\n$") {
			Say("Da fehlt noch ein Leerzeichen zwischen `cd` und `..`!")
		}
		if bash.Test("$(pwd) = \"$ROOT\"") {
			break
		}
	}

	Say("Gut. So, nun brauchen wir ein wenig Magie... *pling*")
	Say("[Der Schrank rumpelt und ächzt]")
	bash.Execute("mkdir -p $ROOT/schrank/magische_tür/tür{1..3}")
	bash.Execute("touch $ROOT/schrank/magische_tür/tür2/rechter_schuh")
	Say("Im Schrank hat sich nun etwas verändert. Geh hinein und such den rechten Schuh.")

	for Prompt() {
		Output()
		if bash.Test("$(pwd) = \"$ROOT/schrank/magische_tür/tür2\"") {
			if Command("^ls\n$") {
				Say("Du hast ihn gefunden! Nun komm wieder zurück!")
				break
			}
		}
	}

	for Prompt() {
		Output()
		if bash.Test("$(pwd) = \"$ROOT\"") {
			break
		}
	}

	Say("Gut! So, das war eine Einführung in `ls` und `cd`. Hier ist das Tutorial erstmal zu Ende! Danke!")
}

func gotoJail() {
	bash.Execute("ROOT=/tmp/nutsh")
	bash.Execute("rm -rf $ROOT")
	bash.Execute("mkdir $ROOT")
	bash.Execute("cd $ROOT")
}
