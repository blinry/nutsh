package main

import (
	. "morr.cc/nutsh.git/dsl"
	"morr.cc/nutsh.git/dsl/bash"
)

func main() {
	Spawn("bash")

	bash.Execute("rm -rf /tmp/nutsh")
	bash.Execute("mkdir /tmp/nutsh")
	bash.Execute("cd /tmp/nutsh")
	bash.Execute("ROOT=/tmp/nutsh")
	bash.Execute("declare -a DIRS=( /tmp /usr )")

	Say("Heute: UNIX-Verzeichnisbaum. Geh mal nach `/` und schau dich um.")

	for prompt() {
		if bash.Test("$(pwd) == \"/\"") && Command("^ls") {
			break
		}
	}

	Say("Nun such dir eins davon aus und sieh dich darin um.")

	for prompt() {
		if lsIn("/tmp") {
			Say("Dies ist die Müllhalde. Sag mal `mount`.")
			for prompt() {
				stayin("/tmp")
				if Command("^mount") {
					Say("Siehste?")
					break
				}
			}
			rm("\\/tmp")
		} else if lsIn("/usr") {
			Say("Hier ist Kram")
			rm("\\/usr")
		} else if Command("^ls") && ! lsIn("/") {
			Say("Hierüber weiß ich nichts")
		}
		if Command("help") {
			Say("Besuche noch folgende Ordner:")
			Say(Query("echo ${DIRS[@]}"))
		}
		if bash.Test("${#DIRS[@]} = 0") {
			break
		}
	}

	Say("Supa!")
}

func lsIn(dir string) bool {
	return bash.Test("$(pwd) == \""+dir+"\"") && Command("^ls")
}

func rm(dir string) {
	bash.Execute("declare -a DIRS=( ${DIRS[@]/"+dir+"/} )")
}

func prompt() bool {
	Prompt()
	bash.Execute("RET=$?")
	if bash.Test("$RET != 0") {
		Say("Da ist wohl was schiefgegangen.")
		if Command("\\.\\.") {
			Say("Wenn du das benutzen möchtest: `alias ..=\"cd ..\"")
		}
	}
	return true
}

func stayin(dir string) {
	if bash.Test("! $(pwd) =~ ^"+dir) {
		Say("Wir sind hier noch nicht fertig! Bitte komm zurück nach `"+dir+"`")
		for prompt() {
			if bash.Test("$(pwd) =~ ^"+dir) {
				break
			}
		}
	}
}
