package carcipher

import "fmt"

func CarstenCipher(message []rune, krypterMed int) []rune {
	return krypter(message, krypterMed)
}

func krypter(message []rune, krypterMed int) []rune {

	kryptertMelding := make([]rune, len(message)) // det er denne som lagrer kryptert/dekryptert streng i typen []rune

	for i := 0; i < len(message); i++ { // looper gjennom hver karakter/tegn i meldingen

		asciiNummer := int(message[i]) // gjør om fra type rune til int

		kryptertMelding[i] = carCipher(asciiNummer, krypterMed, i) // super hemmelig kryptering (tar sikkert minst 2 nanosekund å knekke)

	}

	fmt.Println(string(kryptertMelding))
	return kryptertMelding
}

// the whole point is to not understand this, so you have to figure out on ur own :)
func carCipher(asciiNummer int, krypterMed int, i int) rune {
	// sjekker om koden krypterer med positiv eller negativ int
	if krypterMed > 0 {
		if asciiNummer%2 == 0 {
			asciiNummer *= 2
		}
		return rune(asciiNummer + extraSecurity(i)) // cast til rune
	}
	if krypterMed < 0 {
		asciiNummer -= extraSecurity(i)
		if asciiNummer%2 == 0 {
			asciiNummer /= 2
		}
		return rune(asciiNummer)
	}
	// koden under utføres dersom "krypterMed" == 0
	fmt.Println("Advarsel, du har ikke kryptert meldingen ")
	return rune(asciiNummer)
}

/**
 * Legger til litt ekstra sikkerhet, slik at det blir vanskeliggere å se mønster
 * @Param int sett inn ett satt nummer som blir det samme under en eventuell dekryptering (eks indeksen i en for loop)
 */
func extraSecurity(i int) int {
	if i%2 == 0 {
		return 3
	}
	return 0
}
