The **Nut Shell** is a framework for creating interactive command line tutorials, inspired by text adventures. It can be used to teach system shells, interpreted programming languages, or subtopics, like specific command line tools or concepts.

## Usage

To start a Nut Shell tutorial, use `nutsh run LESSON_DIR`. To test a tutorial for proper function, use `nutsh test LESSON_DIR`.

## Example session

        Do you know the old joke "How do you put an elephant into a fridge?" 

        "Open the fridge, put in the elephant, and close the door". Try it, we 
        delivered a fresh elephant into your kitchen! 

    $ cd kitchen
    $ ls
    elephant fridge/
    $ mv elephant fridge

        [The elephant does not fit into the fridge] 

        Oh, it doesn't seem to be that easy. Can you find out how big the file 
        is? The man page of ls will help you! 

    $ ls
    elephant fridge/
    $ man ls

    [Display of the man page, skipped here]

    $ ls -l
    -rw------- 1 seb users 10485760 27. Okt 22:25 elephant
    drwx------ 2 seb users     4096 27. Okt 22:25 fridge

        Okay, about ten million bytes. ls has the option -h to display that in
        a more comprehensible order of magnitude.

    $ ls -sh
     10M elephant   4.0K fridge

        10 megabytes? Indeed, the fridge isn't that large. 

        We have to make the elephant smaller. There are several ways to 
        compress files - one of the most common under Linux is gzip. 
        Compress the elephant with that program and look at it's size again. 

    $ gzip -f elephant
    $ ls -sh
     3.4M elephant.gz  4.0K fridge

        That sounds much better. 

        Because that went so well, we've delivered a second elephant to the 
        hallway. Get it into the kitchen and and compress it, this time with 
        bzip2. How large is the file this time? 

    $ ls ..
    bedroom/ elephant2 kitchen/ livingroom/
    $ cp ../elephant2 .
    $ bzip2 -f elephant2
    $ ls -sh
    1.9M elephant2.bz2  3.4M elephant.gz  4.0K fridge

        The bzip2 compression algorithm takes more time, but creates smaller 
        files. 

        The two compressed elephants now fit in the fridge comfortably! (Hint: 
        Using wildcards you can move both files at once!) 

    $ mv elephant* fridge

        These commands only compress single files. To compress a folder, you 
        have to combine it and its contents to a single file, you can do that 
        with tar -c -f name_of_the_archive.tar folder. 

        -c stands for "create", -f for the name of the resulting file. tar is
        short for "tape archive" originates from the good old times, when
        magnetic tapes were used to store files. 

        Please combine the fridge to fridge.tar. 

    $ tar -c -f fridge.tar fridge

        The compression commands replace the original file, but tar 
        leaves the source folder intact. Please delete the original fridge. 

    $ rm -rf fridge

        Good, and now compress the so-called "tar ball" with a compression 
        method of your choice. 

    $ gzip fridge.tar

        You can now pick up the compressed fridge with the compressed 
        elephants, put it in your pocket and take it with you. :-) 

        And eventually, we want to get the elephants back. Reverse your steps 
        and put both elephants back into the kitchen. 

        Use the commands gunzip and bunzip2, as well as tar with the arguments
        -x -f name_of_the_archive.tar (-x means "extract"). 

    $ tar xzf fridge.tar.gz
    $ gunzip fridge/elephant.gz

        [The elephant does not fit into the fridge] 

    $ bunzip2 fridge/elephant2.bz2

        [The elephant does not fit into the fridge] 

        Enjoy your meal! When you're ready to complete the lesson, say "done"! 

    $ echo done
    done

## The *nutsh* language

The Nut Shell comes with a domain specific language, that can be used to write tutorial lessons. To give you an idea of *nutsh*'s syntax, here's the source code for the above lesson. First we define some macros:

    // create a folder as an environment for the lesson
    def make_home {
        run(`ROOT="$HOME"/.nutsh`)
        run(`mkdir -p "$ROOT"`)
        run(`cd "$ROOT"`)
        run(`mkdir -p kitchen bedroom livingroom`)
    }

    // return the exit code of the last command
    def exit_code {
        return(run("echo $?"))
    }

    // check a conditional expression using bash's "[[ ]]" syntax
    def test(condition) {
        run(cmd)
        run("[[ "+condition+" ]]")
        return(exit_code == "0")
    }

    // check whether the directory 'd' exists
    def dir(d) {
        return(test("-d \""+d+"\""))
    }

    // check whether the file 'f' exists
    def file(f) {
        return(test("-f \""+f+"\""))
    }

    // was the last command "echo done" or "echo ready"?
    def done {
        return(command =~ `echo\s+(done|ready)`)
    }

And here's the actual lesson:

    make_home

    // create a file that contains 10 MiB of ascending numbers
    run(`seq 1 2000000 | tr -d '\n' | head -c $((1024*1024*10)) >
         "$ROOT/kitchen/elephant"`)

    "Do you know the old joke \"How do you put an elephant into a fridge?\""

    "\"Open the fridge, put in the elephant, and close the door\". Try it, we
    delivered a fresh elephant into your kitchen!"

    // check whether there's an elephant (with suffix 'n') in the fridge
    def elephant_in_fridge(n) {
      return(file(`"$ROOT/kitchen/fridge/elephant"`+n))
    }

    // when there's an elephant in the fridge, show an error and remove it
    def stop_elephant(n) {
      if elephant_in_fridge(n) {
        "[The elephant does not fit into the fridge]"
        run(`mv "$ROOT/kitchen/fridge/elephant`+n+
            `" "$ROOT/kitchen/elephant"`+n)
      }
    }

    prompt {
      if elephant_in_fridge("") {
        stop_elephant("")
        expect("cd kitchen; mv elephant fridge")
        break
      }
    }

    // using a nesting statement, disallow elephants in the fridge for the rest
    // of the tutorial
    stop_elephant(""), stop_elephant("2") {
      "Oh, it doesn't seem to be that easy. Can you find out how big the file
      is? The manpage of `ls` will help you!"

      prompt {
        if output =~ "10M" {
          expect("ls -sh")
          "10 megabytes? Indeed, the fridge isn't that large."
          break
        }
        if output =~ "10240" {
          expect("ls -s")
          "Okay, about ten thousand kilobytes. `ls` has the option `-h` to
          display that in a more comprehensible order of magnitude."
        }
        if output =~ "10485760" {
          expect("ls -l")
          "Okay, about ten million bytes. `ls` has the option `-h` to display
          that in a more comprehensible order of magnitude."
        }
      }

      "We have to make the elephant smaller. There are several ways to
      compress files - one of the most common under Linux is `gzip`. Compress
      the elephant with that program and look at it's size again."

      prompt {
        if file(`"$ROOT/kitchen/elephant.gz"`) {
          // depending on the locale, the decimal mark can be a comma or a dot
          if output =~ `\d[.,]\dM` {
            expect("gzip -f elephant; ls -sh")
            break
          }
        }
      }

      "That sounds much better."

      "Because that went so well, we've delivered a second elephant to the
      hallway. Get it into the kitchen and compress it, this time with
      `bzip2`. How large is the file this time?"

      run(`seq 1 2000000 | tr -d '\n' | head -c $((1024*1024*10)) >
          "$ROOT/elephant2"`)

      prompt {
        if file(`"$ROOT/kitchen/elephant2.bz2"`) {
          if output =~ `\d[.,]\dM` {
            expect("cp ../elephant2 .; bzip2 -f elephant2; ls -sh")
            break
          }
        }
      }

      "The bzip2 compression algorithm takes more time, but creates smaller
      files."

      "The two compressed elephants now fit in the fridge comfortably!
      (Hint: Using wildcards you can move both files at once!)"

      prompt {
        if file(`"$ROOT/kitchen/fridge/elephant.gz"`) &&
           file(`"$ROOT/kitchen/fridge/elephant2.bz2"`) {
          expect(`mv elephant* fridge`)
          break
        }
      }

      "These commands only compress single files. To compress a folder, you
      have to combine it and its contents to a single file, you can do that
      with `tar -c -f name_of_the_archive.tar folder`."

      "`-c` stands for \"create\", `-f` for the name of the resulting file.
      `tar` is short for \"tape archive\" originates from the good old times,
      when magnetic tapes were used to store files."

      "Please combine the fridge to `fridge.tar`."

      prompt {
        if file(`"$ROOT/kitchen/fridge.tar"`) {
          expect("tar -c -f fridge.tar fridge")
          break
        }
      }

      "The compression commands replace the original file, but `tar` leaves
      the source folder intact. Please delete the original fridge."

      prompt {
        if ! dir(`"$ROOT/kitchen/fridge"`) {
          expect("rm -rf fridge")
          break
        }
      }

      "Good, and now compress the so-called \"tar ball\" with a compression
      method of your choice."

      prompt {
        if file(`"$ROOT/kitchen/fridge.tar.gz"`) ||
           file(`"$ROOT/kitchen/fridge.tar.bz2"`) {
          expect("gzip fridge.tar")
          break
        }
      }

      "You can now pick up the compressed fridge with the compressed elephants,
      put it in your pocket and take it with you. :-)"

      "And eventually, we want to get the elephants back. Reverse your steps
      and put both elephants back into the kitchen."

      "Use the commands `gunzip` and `bunzip2`, as well as `tar` with the
      arguments `-x -f name_of_the_archive.tar` (`-x` means \"extract\")."

      prompt {
        if file(`"$ROOT/kitchen/elephant"`) &&
           file(`"$ROOT/kitchen/elephant2"`) {
          expect("tar xzf fridge.tar.gz; gunzip fridge/elephant.gz;
                  bunzip2 fridge/elephant2.bz2")
          break
        }
      }
    }

    "Enjoy your meal! When you're ready to complete the lesson, say \"done\"!"

    prompt {
      if done {
        expect("echo done")
        break
      }
    }
