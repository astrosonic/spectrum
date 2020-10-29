/*
Spectrum v0.1.0 by t0xic0der, mr-tafreshi
Last modified by mr-tafreshi, 13:56 UTC+03:30 17 October 2020
*/

package spectrum

// BACKGROUND COLORING

func BG_BLACK(textobjc string) string { return "\u001b[40m" + textobjc + "\u001b[0m" }

func BG_RED(textobjc string) string { return "\u001b[41m" + textobjc + "\u001b[0m" }

func BG_GREEN(textobjc string) string { return "\u001b[42m" + textobjc + "\u001b[0m" }

func BG_YELLOW(textobjc string) string { return "\u001b[43m" + textobjc + "\u001b[0m" }

func BG_BLUE(textobjc string) string { return "\u001b[44m" + textobjc + "\u001b[0m" }

func BG_MAGENTA(textobjc string) string { return "\u001b[45m" + textobjc + "\u001b[0m" }

func BG_CYAN(textobjc string) string { return "\u001b[46m" + textobjc + "\u001b[0m" }

func BG_WHITE(textobjc string) string { return "\u001b[47m" + textobjc + "\u001b[0m" }

// FOREGROUND COLORING

func FR_BLACK(textobjc string) string { return "\u001b[30m" + textobjc + "\u001b[0m" }

func FR_RED(textobjc string) string { return "\u001b[31m" + textobjc + "\u001b[0m" }

func FR_GREEN(textobjc string) string { return "\u001b[32m" + textobjc + "\u001b[0m" }

func FR_YELLOW(textobjc string) string { return "\u001b[33m" + textobjc + "\u001b[0m" }

func FR_BLUE(textobjc string) string { return "\u001b[34m" + textobjc + "\u001b[0m" }

func FR_MAGENTA(textobjc string) string { return "\u001b[35m" + textobjc + "\u001b[0m" }

func FR_CYAN(textobjc string) string { return "\u001b[36m" + textobjc + "\u001b[0m" }

func FR_WHITE(textobjc string) string { return "\u001b[37m" + textobjc + "\u001b[0m" }

// CUSTOM STYLING

func TX_BOLD(textobjc string) string { return "\u001b[1m" + textobjc + "\u001b[0m" }

func TX_UNDERLINE(textobjc string) string { return "\u001b[4m" + textobjc + "\u001b[0m" }

func TX_INVERTED(textobjc string) string { return "\u001b[7m" + textobjc + "\u001b[0m" }

func HD_SUCCESS(textobjc string) string { return FR_GREEN(TX_BOLD("[SUCCESS] " + textobjc)) }

func HD_FAILURE(textobjc string) string { return FR_RED(TX_BOLD("[FAILURE] " + textobjc)) }

func HD_DETAILS(textobjc string) string { return FR_BLUE(TX_BOLD("[DETAILS] " + textobjc)) }

func HD_WARNING(textobjc string) string { return FR_YELLOW(TX_BOLD("[WARNING] " + textobjc)) }
