package main 

import (
    "fmt"
    "math/rand"
    "time"
    "strconv"
)

func random_number_generator(min, max int) int {
    return rand.Intn(max-min+1) + min
}

func welcome() {
	fmt.Println("Welcome to the game 'Guess the secret number if you can'")
	fmt.Println("Please enjoy the game while you can! Cause it's game over for you once you fail")

	fmt.Println()

	frames := []string{"---/---", "---|---", "---\\---", "---/---"}
	duration := 4 * time.Second  // Total time for animation
	interval := 200 * time.Millisecond // Speed of animation

	startTime := time.Now()

	for time.Since(startTime) < duration {
		for _, frame := range frames {
			if time.Since(startTime) >= duration {
				break
			}
			fmt.Print("\r" + frame)
			time.Sleep(interval)
		}
	}

	fmt.Print("\r") 
	fmt.Println("Just kidding... You can always retry and play the game until you win!!")
}

func input() (string, int) {
	var username string
	fmt.Print("Create a username for the game (your in-game name/what we will call you in the game): ")
	fmt.Scanln(&username)

	fmt.Println("\nChoose the difficulty level of the game:")
	fmt.Println("1. Casual Level")
	fmt.Println("2. Standard Level")
	fmt.Println("3. Challenging Level 1")
	fmt.Println("4. Challenging Level 2")
	fmt.Println("5. Extreme Level")
	fmt.Println("6. Impossible Level")

	var difficulty int
	for {
		fmt.Print("\nEnter a number between (1-6) to choose a difficulty level: ")
		var input string
		fmt.Scanln(&input)

		num, err := strconv.Atoi(input)
		if err != nil || num < 1 || num > 6 {
			fmt.Println("Invalid choice! Please enter a number between 1 and 6.")
		} else {
			difficulty = num
			break
		}
	}

	return username, difficulty
}


func getOrdinalSuffix(n int) string {
	if n%10 == 1 && n%100 != 11 {
		return "st"
	} else if n%10 == 2 && n%100 != 12 {
		return "nd"
	} else if n%10 == 3 && n%100 != 13 {
		return "rd"
	}
	return "th"
}

/*                                                                                                
        CCCCCCCCCCCCC                                                                      lllllll 
     CCC::::::::::::C                                                                      l:::::l 
   CC:::::::::::::::C                                                                      l:::::l 
  C:::::CCCCCCCC::::C                                                                      l:::::l 
 C:::::C       CCCCCC  aaaaaaaaaaaaa      ssssssssss   uuuuuu    uuuuuu    aaaaaaaaaaaaa    l::::l 
C:::::C                a::::::::::::a   ss::::::::::s  u::::u    u::::u    a::::::::::::a   l::::l 
C:::::C                aaaaaaaaa:::::ass:::::::::::::s u::::u    u::::u    aaaaaaaaa:::::a  l::::l 
C:::::C                         a::::as::::::ssss:::::su::::u    u::::u             a::::a  l::::l 
C:::::C                  aaaaaaa:::::a s:::::s  ssssss u::::u    u::::u      aaaaaaa:::::a  l::::l 
C:::::C                aa::::::::::::a   s::::::s      u::::u    u::::u    aa::::::::::::a  l::::l 
C:::::C               a::::aaaa::::::a      s::::::s   u::::u    u::::u   a::::aaaa::::::a  l::::l 
 C:::::C       CCCCCCa::::a    a:::::assssss   s:::::s u:::::uuuu:::::u  a::::a    a:::::a  l::::l 
  C:::::CCCCCCCC::::Ca::::a    a:::::as:::::ssss::::::su:::::::::::::::uua::::a    a:::::a l::::::l
   CC:::::::::::::::Ca:::::aaaa::::::as::::::::::::::s  u:::::::::::::::ua:::::aaaa::::::a l::::::l
     CCC::::::::::::C a::::::::::aa:::as:::::::::::ss    uu::::::::uu:::u a::::::::::aa:::al::::::l
        CCCCCCCCCCCCC  aaaaaaaaaa  aaaa sssssssssss        uuuuuuuu  uuuu  aaaaaaaaaa  aaaallllllll
                                                                                                   
*/



func casual_level() int {
	secretNumber := random_number_generator(1, 50)
	attempts := 8 

	fmt.Println("\nWelcome to the Casual Level!")
	fmt.Println("I have selected a number between 1 and 50.")
	fmt.Println("You have 8 attempts to guess it correctly.")
	fmt.Println("Let's begin!")
    fmt.Println()

	for i := 1; i <= attempts; i++ {
		fmt.Printf("Attempt %d: Enter your guess: ", i)
		var input string
		fmt.Scanln(&input)

		// Convert input to an integer
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input! Please enter a valid number.")
			i-- // Don't count invalid inputs as attempts
			continue
		}


		if guess > secretNumber {
			fmt.Println("Too high!")
		} else if guess < secretNumber {
			fmt.Println("Too low!")
		} else {
			fmt.Printf("Good job! You guessed it correctly.\n")
			fmt.Printf("You guessed it at your %d%s turn.\n", i, getOrdinalSuffix(i))
			return i
		}
	}

	// If user runs out of attempts
	fmt.Printf("\nGame Over! The correct number was %d.\n", secretNumber)
    return 0
}

func casual_level_scoring(username string, attempt int) {
	totalScore := 0
	maxAttempts := 8

	if attempt > 0 { // If the user won
		totalScore = 50 + (maxAttempts-attempt)*10
	}

	// Print the final score
	fmt.Printf("\nUsername: %s\n", username)
	fmt.Printf("Total Score: %d\n", totalScore)
}

/*

                                                                                                                                                                                                                                           dddddddd                                                 dddddddd
   SSSSSSSSSSSSSSS      tttt                                                         d::::::d                                                 d::::::d
 SS:::::::::::::::S  ttt:::t                                                         d::::::d                                                 d::::::d
S:::::SSSSSS::::::S  t:::::t                                                         d::::::d                                                 d::::::d
S:::::S     SSSSSSS  t:::::t                                                         d:::::d                                                  d:::::d 
S:::::S        ttttttt:::::ttttttt      aaaaaaaaaaaaa  nnnn  nnnnnnnn        ddddddddd:::::d   aaaaaaaaaaaaa  rrrrr   rrrrrrrrr       ddddddddd:::::d 
S:::::S        t:::::::::::::::::t      a::::::::::::a n:::nn::::::::nn    dd::::::::::::::d   a::::::::::::a r::::rrr:::::::::r    dd::::::::::::::d 
 S::::SSSS     t:::::::::::::::::t      aaaaaaaaa:::::an::::::::::::::nn  d::::::::::::::::d   aaaaaaaaa:::::ar:::::::::::::::::r  d::::::::::::::::d 
  SS::::::SSSSStttttt:::::::tttttt               a::::ann:::::::::::::::nd:::::::ddddd:::::d            a::::arr::::::rrrrr::::::rd:::::::ddddd:::::d 
    SSS::::::::SS    t:::::t              aaaaaaa:::::a  n:::::nnnn:::::nd::::::d    d:::::d     aaaaaaa:::::a r:::::r     r:::::rd::::::d    d:::::d 
       SSSSSS::::S   t:::::t            aa::::::::::::a  n::::n    n::::nd:::::d     d:::::d   aa::::::::::::a r:::::r     rrrrrrrd:::::d     d:::::d 
            S:::::S  t:::::t           a::::aaaa::::::a  n::::n    n::::nd:::::d     d:::::d  a::::aaaa::::::a r:::::r            d:::::d     d:::::d 
            S:::::S  t:::::t    tttttta::::a    a:::::a  n::::n    n::::nd:::::d     d:::::d a::::a    a:::::a r:::::r            d:::::d     d:::::d 
SSSSSSS     S:::::S  t::::::tttt:::::ta::::a    a:::::a  n::::n    n::::nd::::::ddddd::::::dda::::a    a:::::a r:::::r            d::::::ddddd::::::dd
S::::::SSSSSS:::::S  tt::::::::::::::ta:::::aaaa::::::a  n::::n    n::::n d:::::::::::::::::da:::::aaaa::::::a r:::::r             d:::::::::::::::::d
S:::::::::::::::SS     tt:::::::::::tt a::::::::::aa:::a n::::n    n::::n  d:::::::::ddd::::d a::::::::::aa:::ar:::::r              d:::::::::ddd::::d
 SSSSSSSSSSSSSSS         ttttttttttt    aaaaaaaaaa  aaaa nnnnnn    nnnnnn   ddddddddd   ddddd  aaaaaaaaaa  aaaarrrrrrr               ddddddddd   ddddd
                                                                                                                                                      
                                                                                                                                                      
*/




func standard_level() int {
	secretNumber := random_number_generator(1, 100)
	attempts := 6 

	fmt.Println("\nWelcome to the Standard Level!")
	fmt.Println("I have selected a number between 1 and 100.")
	fmt.Println("You have 6 attempts to guess it correctly.")
	fmt.Println("Let's begin!")
    fmt.Println()

	for i := 1; i <= attempts; i++ {
		fmt.Printf("Attempt %d: Enter your guess: ", i)
		var input string
		fmt.Scanln(&input)

		// Convert input to an integer
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input! Please enter a valid number.")
			i-- // Don't count invalid inputs as attempts
			continue
		}


		if guess > secretNumber {
			fmt.Println("Too high!")
		} else if guess < secretNumber {
			fmt.Println("Too low!")
		} else {
			fmt.Printf("Good job! You guessed it correctly.\n")
			fmt.Printf("You guessed it at your %d%s turn.\n", i, getOrdinalSuffix(i))
			return i
		}
	}

	// If user runs out of attempts
	fmt.Printf("\nGame Over! The correct number was %d.\n", secretNumber)
    return 0
}


func standard_level_scoring(username string, attempt int) {
	totalScore := 0
	maxAttempts := 6

	if attempt > 0 { // If the user won
		totalScore = 100 + (maxAttempts-attempt)*50
	}

	// Print the final score
	fmt.Printf("\nUsername: %s\n", username)
	fmt.Printf("Total Score: %d\n", totalScore)
}

/*


                                                                                                                                                                   
                                                                                                                                                                   
        CCCCCCCCCCCCChhhhhhh                               lllllll lllllll                                                                               1111111   
     CCC::::::::::::Ch:::::h                               l:::::l l:::::l                                                                              1::::::1   
   CC:::::::::::::::Ch:::::h                               l:::::l l:::::l                                                                             1:::::::1   
  C:::::CCCCCCCC::::Ch:::::h                               l:::::l l:::::l                                                                             111:::::1   
 C:::::C       CCCCCC h::::h hhhhh         aaaaaaaaaaaaa    l::::l  l::::l     eeeeeeeeeeee    nnnn  nnnnnnnn       ggggggggg   ggggg    eeeeeeeeeeee     1::::1   
C:::::C               h::::hh:::::hhh      a::::::::::::a   l::::l  l::::l   ee::::::::::::ee  n:::nn::::::::nn    g:::::::::ggg::::g  ee::::::::::::ee   1::::1   
C:::::C               h::::::::::::::hh    aaaaaaaaa:::::a  l::::l  l::::l  e::::::eeeee:::::een::::::::::::::nn  g:::::::::::::::::g e::::::eeeee:::::ee 1::::1   
C:::::C               h:::::::hhh::::::h            a::::a  l::::l  l::::l e::::::e     e:::::enn:::::::::::::::ng::::::ggggg::::::gge::::::e     e:::::e 1::::l   
C:::::C               h::::::h   h::::::h    aaaaaaa:::::a  l::::l  l::::l e:::::::eeeee::::::e  n:::::nnnn:::::ng:::::g     g:::::g e:::::::eeeee::::::e 1::::l   
C:::::C               h:::::h     h:::::h  aa::::::::::::a  l::::l  l::::l e:::::::::::::::::e   n::::n    n::::ng:::::g     g:::::g e:::::::::::::::::e  1::::l   
C:::::C               h:::::h     h:::::h a::::aaaa::::::a  l::::l  l::::l e::::::eeeeeeeeeee    n::::n    n::::ng:::::g     g:::::g e::::::eeeeeeeeeee   1::::l   
 C:::::C       CCCCCC h:::::h     h:::::ha::::a    a:::::a  l::::l  l::::l e:::::::e             n::::n    n::::ng::::::g    g:::::g e:::::::e            1::::l   
  C:::::CCCCCCCC::::C h:::::h     h:::::ha::::a    a:::::a l::::::ll::::::le::::::::e            n::::n    n::::ng:::::::ggggg:::::g e::::::::e        111::::::111
   CC:::::::::::::::C h:::::h     h:::::ha:::::aaaa::::::a l::::::ll::::::l e::::::::eeeeeeee    n::::n    n::::n g::::::::::::::::g  e::::::::eeeeeeee1::::::::::1
     CCC::::::::::::C h:::::h     h:::::h a::::::::::aa:::al::::::ll::::::l  ee:::::::::::::e    n::::n    n::::n  gg::::::::::::::g   ee:::::::::::::e1::::::::::1
        CCCCCCCCCCCCC hhhhhhh     hhhhhhh  aaaaaaaaaa  aaaallllllllllllllll    eeeeeeeeeeeeee    nnnnnn    nnnnnn    gggggggg::::::g     eeeeeeeeeeeeee111111111111
                                                                                                                             g:::::g                               
                                                                                                                 gggggg      g:::::g                               
                                                                                                                 g:::::gg   gg:::::g                               
                                                                                                                  g::::::ggg:::::::g                               
                                                                                                                   gg:::::::::::::g                                
                                                                                                                     ggg::::::ggg                                  
                                                                                                                        gggggg                                     


*/
func challenging_level_1() int {
	secretNumber := random_number_generator(1, 100)
	attempts := 5 

	fmt.Println("\nWelcome to the Challenging Level 1!")
	fmt.Println("I have selected a number between 1 and 100.")
	fmt.Println("You have 5 attempts to guess it correctly.")
	fmt.Println("Let's begin!")
    fmt.Println()

	for i := 1; i <= attempts; i++ {
		fmt.Printf("Attempt %d: Enter your guess: ", i)
		var input string
		fmt.Scanln(&input)

		// Convert input to an integer
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input! Please enter a valid number.")
			i-- // Don't count invalid inputs as attempts
			continue
		}


		if guess > secretNumber {
			fmt.Println("Too high!")
		} else if guess < secretNumber {
			fmt.Println("Too low!")
		} else {
			fmt.Printf("Good job! You guessed it correctly.\n")
			fmt.Printf("You guessed it at your %d%s turn.\n", i, getOrdinalSuffix(i))
			return i
		}
	}

	// If user runs out of attempts
	fmt.Printf("\nGame Over! The correct number was %d.\n", secretNumber)
    return 0
}


func challenging_level_1_scoring(username string, attempt int) {
	totalScore := 0
	maxAttempts := 5

	if attempt > 0 { // If the user won
		totalScore = 300 + (maxAttempts-attempt)*100
	}

	// Print the final score
	fmt.Printf("\nUsername: %s\n", username)
	fmt.Printf("Total Score: %d\n", totalScore)
}
/*


                                                                                                                                                                           
                                                                                                                                                                           
        CCCCCCCCCCCCChhhhhhh                               lllllll lllllll                                                                              222222222222222    
     CCC::::::::::::Ch:::::h                               l:::::l l:::::l                                                                             2:::::::::::::::22  
   CC:::::::::::::::Ch:::::h                               l:::::l l:::::l                                                                             2::::::222222:::::2 
  C:::::CCCCCCCC::::Ch:::::h                               l:::::l l:::::l                                                                             2222222     2:::::2 
 C:::::C       CCCCCC h::::h hhhhh         aaaaaaaaaaaaa    l::::l  l::::l     eeeeeeeeeeee    nnnn  nnnnnnnn       ggggggggg   ggggg    eeeeeeeeeeee              2:::::2 
C:::::C               h::::hh:::::hhh      a::::::::::::a   l::::l  l::::l   ee::::::::::::ee  n:::nn::::::::nn    g:::::::::ggg::::g  ee::::::::::::ee            2:::::2 
C:::::C               h::::::::::::::hh    aaaaaaaaa:::::a  l::::l  l::::l  e::::::eeeee:::::een::::::::::::::nn  g:::::::::::::::::g e::::::eeeee:::::ee       2222::::2  
C:::::C               h:::::::hhh::::::h            a::::a  l::::l  l::::l e::::::e     e:::::enn:::::::::::::::ng::::::ggggg::::::gge::::::e     e:::::e  22222::::::22   
C:::::C               h::::::h   h::::::h    aaaaaaa:::::a  l::::l  l::::l e:::::::eeeee::::::e  n:::::nnnn:::::ng:::::g     g:::::g e:::::::eeeee::::::e22::::::::222     
C:::::C               h:::::h     h:::::h  aa::::::::::::a  l::::l  l::::l e:::::::::::::::::e   n::::n    n::::ng:::::g     g:::::g e:::::::::::::::::e2:::::22222        
C:::::C               h:::::h     h:::::h a::::aaaa::::::a  l::::l  l::::l e::::::eeeeeeeeeee    n::::n    n::::ng:::::g     g:::::g e::::::eeeeeeeeeee2:::::2             
 C:::::C       CCCCCC h:::::h     h:::::ha::::a    a:::::a  l::::l  l::::l e:::::::e             n::::n    n::::ng::::::g    g:::::g e:::::::e         2:::::2             
  C:::::CCCCCCCC::::C h:::::h     h:::::ha::::a    a:::::a l::::::ll::::::le::::::::e            n::::n    n::::ng:::::::ggggg:::::g e::::::::e        2:::::2       222222
   CC:::::::::::::::C h:::::h     h:::::ha:::::aaaa::::::a l::::::ll::::::l e::::::::eeeeeeee    n::::n    n::::n g::::::::::::::::g  e::::::::eeeeeeee2::::::2222222:::::2
     CCC::::::::::::C h:::::h     h:::::h a::::::::::aa:::al::::::ll::::::l  ee:::::::::::::e    n::::n    n::::n  gg::::::::::::::g   ee:::::::::::::e2::::::::::::::::::2
        CCCCCCCCCCCCC hhhhhhh     hhhhhhh  aaaaaaaaaa  aaaallllllllllllllll    eeeeeeeeeeeeee    nnnnnn    nnnnnn    gggggggg::::::g     eeeeeeeeeeeeee22222222222222222222
                                                                                                                             g:::::g                                       
                                                                                                                 gggggg      g:::::g                                       
                                                                                                                 g:::::gg   gg:::::g                                       
                                                                                                                  g::::::ggg:::::::g                                       
                                                                                                                   gg:::::::::::::g                                        
                                                                                                                     ggg::::::ggg                                          
                                                                                                                        gggggg                                             


*/
func challenging_level_2() int {
	secretNumber := random_number_generator(1, 100)
	attempts := 4 

	fmt.Println("\nWelcome to the Challenging Level 2!")
	fmt.Println("I have selected a number between 1 and 100.")
	fmt.Println("You have 4 attempts to guess it correctly.")
	fmt.Println("Let's begin!")
    fmt.Println()

	for i := 1; i <= attempts; i++ {
		fmt.Printf("Attempt %d: Enter your guess: ", i)
		var input string
		fmt.Scanln(&input)

		// Convert input to an integer
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input! Please enter a valid number.")
			i-- // Don't count invalid inputs as attempts
			continue
		}


		if guess > secretNumber {
			fmt.Println("Too high!")
		} else if guess < secretNumber {
			fmt.Println("Too low!")
		} else {
			fmt.Printf("Good job! You guessed it correctly.\n")
			fmt.Printf("You guessed it at your %d%s turn.\n", i, getOrdinalSuffix(i))
			return i
		}
	}

	// If user runs out of attempts
	fmt.Printf("\nGame Over! The correct number was %d.\n", secretNumber)
    return 0
}


func challenging_level_2_scoring(username string, attempt int) {
	totalScore := 0
	maxAttempts := 4

	if attempt > 0 { // If the user won
		totalScore = 500 + (maxAttempts-attempt)*200
	}

	// Print the final score
	fmt.Printf("\nUsername: %s\n", username)
	fmt.Printf("Total Score: %d\n", totalScore)
}
/*


                                                                                                                                                    
                                                                                                                                                    
EEEEEEEEEEEEEEEEEEEEEE                             tttt                                                                                             
E::::::::::::::::::::E                          ttt:::t                                                                                             
E::::::::::::::::::::E                          t:::::t                                                                                             
EE::::::EEEEEEEEE::::E                          t:::::t                                                                                             
  E:::::E       EEEEEExxxxxxx      xxxxxxxttttttt:::::ttttttt   rrrrr   rrrrrrrrr       eeeeeeeeeeee       mmmmmmm    mmmmmmm       eeeeeeeeeeee    
  E:::::E              x:::::x    x:::::x t:::::::::::::::::t   r::::rrr:::::::::r    ee::::::::::::ee   mm:::::::m  m:::::::mm   ee::::::::::::ee  
  E::::::EEEEEEEEEE     x:::::x  x:::::x  t:::::::::::::::::t   r:::::::::::::::::r  e::::::eeeee:::::eem::::::::::mm::::::::::m e::::::eeeee:::::ee
  E:::::::::::::::E      x:::::xx:::::x   tttttt:::::::tttttt   rr::::::rrrrr::::::re::::::e     e:::::em::::::::::::::::::::::me::::::e     e:::::e
  E:::::::::::::::E       x::::::::::x          t:::::t          r:::::r     r:::::re:::::::eeeee::::::em:::::mmm::::::mmm:::::me:::::::eeeee::::::e
  E::::::EEEEEEEEEE        x::::::::x           t:::::t          r:::::r     rrrrrrre:::::::::::::::::e m::::m   m::::m   m::::me:::::::::::::::::e 
  E:::::E                  x::::::::x           t:::::t          r:::::r            e::::::eeeeeeeeeee  m::::m   m::::m   m::::me::::::eeeeeeeeeee  
  E:::::E       EEEEEE    x::::::::::x          t:::::t    ttttttr:::::r            e:::::::e           m::::m   m::::m   m::::me:::::::e           
EE::::::EEEEEEEE:::::E   x:::::xx:::::x         t::::::tttt:::::tr:::::r            e::::::::e          m::::m   m::::m   m::::me::::::::e          
E::::::::::::::::::::E  x:::::x  x:::::x        tt::::::::::::::tr:::::r             e::::::::eeeeeeee  m::::m   m::::m   m::::m e::::::::eeeeeeee  
E::::::::::::::::::::E x:::::x    x:::::x         tt:::::::::::ttr:::::r              ee:::::::::::::e  m::::m   m::::m   m::::m  ee:::::::::::::e  
EEEEEEEEEEEEEEEEEEEEEExxxxxxx      xxxxxxx          ttttttttttt  rrrrrrr                eeeeeeeeeeeeee  mmmmmm   mmmmmm   mmmmmm    eeeeeeeeeeeeee  
                                                                                                                                                    
                                                                                                                                                    
                                                                                                                                                    
                                                                                                                                                    
                                                                                                                                                    
                                                                                                                                                    
                                                                                                                                                    


*/
func extreme_level() int {
	secretNumber := random_number_generator(1, 500)
	attempts := 8 
	fmt.Println("\nWelcome to the Extreme Level!")
	fmt.Println("I have selected a number between 1 and 500.")
	fmt.Println("You have 8 attempts to guess it correctly.")
	fmt.Println("Let's begin!")
    fmt.Println()

	for i := 1; i <= attempts; i++ {
		fmt.Printf("Attempt %d: Enter your guess: ", i)
		var input string
		fmt.Scanln(&input)

		// Convert input to an integer
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input! Please enter a valid number.")
			i-- // Don't count invalid inputs as attempts
			continue
		}


		if guess > secretNumber {
			fmt.Println("Too high!")
		} else if guess < secretNumber {
			fmt.Println("Too low!")
		} else {
			fmt.Printf("Good job! You guessed it correctly.\n")
			fmt.Printf("You guessed it at your %d%s turn.\n", i, getOrdinalSuffix(i))
			return i
		}
	}

	// If user runs out of attempts
	fmt.Printf("\nGame Over! The correct number was %d.\n", secretNumber)
    return 0
}


func extreme_level_scoring(username string, attempt int) {
	totalScore := 0
	maxAttempts := 8

	if attempt > 0 { // If the user won
		totalScore = 1000 + (maxAttempts-attempt)*500
	}

	// Print the final score
	fmt.Printf("\nUsername: %s\n", username)
	fmt.Printf("Total Score: %d\n", totalScore)
}
/*


                                                                                                                                                                
                                                                                                                bbbbbbbb                                        
IIIIIIIIII                                                                                                 iiii b::::::b            lllllll                     
I::::::::I                                                                                                i::::ib::::::b            l:::::l                     
I::::::::I                                                                                                 iiii b::::::b            l:::::l                     
II::::::II                                                                                                       b:::::b            l:::::l                     
  I::::I     mmmmmmm    mmmmmmm   ppppp   ppppppppp      ooooooooooo       ssssssssss       ssssssssss   iiiiiii b:::::bbbbbbbbb     l::::l     eeeeeeeeeeee    
  I::::I   mm:::::::m  m:::::::mm p::::ppp:::::::::p   oo:::::::::::oo   ss::::::::::s    ss::::::::::s  i:::::i b::::::::::::::bb   l::::l   ee::::::::::::ee  
  I::::I  m::::::::::mm::::::::::mp:::::::::::::::::p o:::::::::::::::oss:::::::::::::s ss:::::::::::::s  i::::i b::::::::::::::::b  l::::l  e::::::eeeee:::::ee
  I::::I  m::::::::::::::::::::::mpp::::::ppppp::::::po:::::ooooo:::::os::::::ssss:::::ss::::::ssss:::::s i::::i b:::::bbbbb:::::::b l::::l e::::::e     e:::::e
  I::::I  m:::::mmm::::::mmm:::::m p:::::p     p:::::po::::o     o::::o s:::::s  ssssss  s:::::s  ssssss  i::::i b:::::b    b::::::b l::::l e:::::::eeeee::::::e
  I::::I  m::::m   m::::m   m::::m p:::::p     p:::::po::::o     o::::o   s::::::s         s::::::s       i::::i b:::::b     b:::::b l::::l e:::::::::::::::::e 
  I::::I  m::::m   m::::m   m::::m p:::::p     p:::::po::::o     o::::o      s::::::s         s::::::s    i::::i b:::::b     b:::::b l::::l e::::::eeeeeeeeeee  
  I::::I  m::::m   m::::m   m::::m p:::::p    p::::::po::::o     o::::ossssss   s:::::s ssssss   s:::::s  i::::i b:::::b     b:::::b l::::l e:::::::e           
II::::::IIm::::m   m::::m   m::::m p:::::ppppp:::::::po:::::ooooo:::::os:::::ssss::::::ss:::::ssss::::::si::::::ib:::::bbbbbb::::::bl::::::le::::::::e          
I::::::::Im::::m   m::::m   m::::m p::::::::::::::::p o:::::::::::::::os::::::::::::::s s::::::::::::::s i::::::ib::::::::::::::::b l::::::l e::::::::eeeeeeee  
I::::::::Im::::m   m::::m   m::::m p::::::::::::::pp   oo:::::::::::oo  s:::::::::::ss   s:::::::::::ss  i::::::ib:::::::::::::::b  l::::::l  ee:::::::::::::e  
IIIIIIIIIImmmmmm   mmmmmm   mmmmmm p::::::pppppppp       ooooooooooo     sssssssssss      sssssssssss    iiiiiiiibbbbbbbbbbbbbbbb   llllllll    eeeeeeeeeeeeee  
                                   p:::::p                                                                                                                      
                                   p:::::p                                                                                                                      
                                  p:::::::p                                                                                                                     
                                  p:::::::p                                                                                                                     
                                  p:::::::p                                                                                                                     
                                  ppppppppp                                                                                                                     
                                                                                                                                                                


*/
func impossible_level() int {
	secretNumber := random_number_generator(1, 1000)
	attempts := 8 

	fmt.Println("\nWelcome to the Impossible Level!")
	fmt.Println("I have selected a number between 1 and 1000.")
	fmt.Println("You have 8 attempts to guess it correctly.")
	fmt.Println("Let's begin!")
    fmt.Println()

	for i := 1; i <= attempts; i++ {
		fmt.Printf("Attempt %d: Enter your guess: ", i)
		var input string
		fmt.Scanln(&input)

		// Convert input to an integer
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input! Please enter a valid number.")
			i-- // Don't count invalid inputs as attempts
			continue
		}


		if guess > secretNumber {
			fmt.Println("Too high!")
		} else if guess < secretNumber {
			fmt.Println("Too low!")
		} else {
			fmt.Printf("Good job! You guessed it correctly.\n")
			fmt.Printf("You guessed it at your %d%s turn.\n", i, getOrdinalSuffix(i))
			return i
		}
	}

	// If user runs out of attempts
	fmt.Printf("\nGame Over! The correct number was %d.\n", secretNumber)
    return 0
}


func impossible_level_scoring(username string, attempt int) {
	totalScore := 0
	maxAttempts := 8

	if attempt > 0 { // If the user won
		totalScore = 3000 + (maxAttempts-attempt)*750
	}

	// Print the final score
	fmt.Printf("\nUsername: %s\n", username)
	fmt.Printf("Total Score: %d\n", totalScore)
}




func main() {
    welcome()
    username, difficulty := input()

    if (difficulty == 1) {
        level_1 := casual_level()
        casual_level_scoring(username, level_1)
    } else if (difficulty == 2) {
        level_2 := standard_level()
        standard_level_scoring(username, level_2)
    } else if (difficulty == 3) {
        level_3 := challenging_level_1()
        challenging_level_1_scoring(username, level_3)
    } else if (difficulty == 4) {
        level_4 := challenging_level_2()
        challenging_level_2_scoring(username, level_4)
    } else if (difficulty == 5) {
        level_5 := extreme_level()
        extreme_level_scoring(username, level_5)
    } else {
        level_6 := impossible_level()
        impossible_level_scoring(username, level_6)
    }
}
