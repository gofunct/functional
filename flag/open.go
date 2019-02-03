package flag

/*

func Open(fs *flag.FlagSet) {
	var (
		fruit    = flaggers.Enum{Choices: []string{"apple", "banana"}}
		urls     flaggers.URLs
		settings flaggers.AssignmentsMap
	)
	fs.Var(&fruit, "fruit", "a fruit")
	fs.Var(&urls, "url", "a URL")
	fs.Var(&settings, "set", "set key=value")
	fs.Parse([]string{
		"-fruit", "apple",
		"-url", "https://github.com/sgreben/flagvar",
		"-set", "hello=world",
	})

	fmt.Println("fruit:", fruit.Value)
	fmt.Println("urls:", urls.Values)
	for key, value := range settings.Values {
		fmt.Printf("settings: '%s' is set to '%s'\n", key, value)
	}
}
 */