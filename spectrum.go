/*
Spectrum v0.1.0 by t0xic0der
Last modified by Akashdeep Dhar, 13:05 UTC+05:30 15 October 2020
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

// CLI

func CLI_SUCCESS(textobjc string) string { return TX_BOLD(FR_GREEN("[SUCCESS] " + textobjc)) }

func CLI_WARNING(textobjc string) string { return TX_BOLD(FR_YELLOW("[WARNING] " + textobjc)) }

func CLI_FAILURE(textobjc string) string { return TX_BOLD(FR_RED("[FAILURE] " + textobjc)) }

func CLI_DETAILS(textobjc string) string { return TX_BOLD(FR_BLUE("[DETAILS] " + textobjc)) }
