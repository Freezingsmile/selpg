package main

import (
	"fmt"
	"io"
	"os"
  "bufio"
	"flag"
	"os/exec"
)

var progname string

type Selpg_args struct {
	start_page int
	end_page int
	in_filename string
	page_len int
	page_type int
	print_dest string
}

func errorExit(reason string) {
	fmt.Fprintln(os.Stderr, reason)
	os.Exit(1)
}

func process_args(sa *Selpg_args) {
	/* check the command-line arguments for validity */
	/* Not enough args, minimum command is "selpg -sstartpage -eend_page"  */
	if sa.start_page == -1 || sa.end_page == -1 {
		fmt.Fprintf(os.Stderr, "%s: not enough arguments\n", progname)
		if sa.end_page != -1 {
			fmt.Fprintf(os.Stderr, "%s: 1st arg should be -sstart_page\n", progname)
		} else {
			fmt.Fprintf(os.Stderr, "%s: 2nd arg should be -eend_page\n", progname)
		}
		flag.Usage()
		os.Exit(1)
	}
	/* handle invalid page rage */
	if sa.start_page < 0 || sa.start_page > sa.end_page || sa.end_page < 0 {
		fmt.Fprintf(os.Stderr, "The range of the page is invalid")
		flag.Usage()
		os.Exit(2)
	}
}

func process_input(sa *Selpg_args) {
	if flag.NArg() < 0 {
		scanner := bufio.NewScanner(os.Stdin)
		counter := 0
		response := ""
		for scanner.Scan() {
			line := scanner.Text()
			line += "\n"
			positions := counter/sa.page_len
			if positions <= sa.end_page && positions >= sa.start_page {
				response += line
			}
			counter++
		}
		fmt.Printf("%s", response)
	} else {
		sa.in_filename = flag.Arg(0)
		fname, err := os.Open(sa.in_filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}
		reader := bufio.NewReader(fname)
		counter := 0
		for {
			line, _, err := reader.ReadLine()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println(err)
				os.Exit(4)
			}
			positions := counter/sa.page_len
			if positions <= sa.end_page && positions >= sa.start_page {
				fmt.Println(string(line))
			}
			counter++
		}
	}
	cmd := exec.Command("cat", "-n")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	if sa.print_dest != "" {
		stdin.Close()
		cmd.Stdout = os.Stdout;
		cmd.Start()
	}
}

func main() {
	progname = os.Args[0]
	var sa Selpg_args
	flag.Usage = usage
	flag.IntVar(&sa.start_page, "s", -1, "sp")
	flag.IntVar(&sa.end_page, "e", -1, "ep")
	flag.IntVar(&sa.page_len, "l", 72, "length/p")
	flag.IntVar(&sa.page_type, "f", 0, "line/form-feed")
	print_dest := flag.String("d", "", "print dest")
	flag.Parse()
	if *print_dest  != "" {
		sa.print_dest = *print_dest
	}
	process_args(&sa)
	process_input(&sa)
}

