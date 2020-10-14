/*
Spectrum v0.0.1 by t0xic0der
Last modified by Akashdeep Dhar, 00:05 UTC+05:30 15 October 2020
*/

package spectrum

// BACKGROUND + BOLD

func BackBlkBold(textobjc string) string {
	return string("\u001b[40;1m") + textobjc + string("\u001b[0m")
}

func BackRedBold(textobjc string) string {
	return string("\u001b[41;1m") + textobjc + string("\u001b[0m")
}

func BackGrnBold(textobjc string) string {
	return string("\u001b[42;1m") + textobjc + string("\u001b[0m")
}

func BackYlwBold(textobjc string) string {
	return string("\u001b[43;1m") + textobjc + string("\u001b[0m")
}

func BackBluBold(textobjc string) string {
	return string("\u001b[44;1m") + textobjc + string("\u001b[0m")
}

func BackMgtBold(textobjc string) string {
	return string("\u001b[45;1m") + textobjc + string("\u001b[0m")
}

func BackCynBold(textobjc string) string {
	return string("\u001b[46;1m") + textobjc + string("\u001b[0m")
}

func BackWitBold(textobjc string) string {
	return string("\u001b[47;1m") + textobjc + string("\u001b[0m")
}

// BACKGROUND + REGULAR

func BackBlkRglr(textobjc string) string {
	return string("\u001b[40m") + textobjc + string("\u001b[0m")
}

func BackRedRglr(textobjc string) string {
	return string("\u001b[41m") + textobjc + string("\u001b[0m")
}

func BackGrnRglr(textobjc string) string {
	return string("\u001b[42m") + textobjc + string("\u001b[0m")
}

func BackYlwRglr(textobjc string) string {
	return string("\u001b[43m") + textobjc + string("\u001b[0m")
}

func BackBluRglr(textobjc string) string {
	return string("\u001b[44m") + textobjc + string("\u001b[0m")
}

func BackMgtRglr(textobjc string) string {
	return string("\u001b[45m") + textobjc + string("\u001b[0m")
}

func BackCynRglr(textobjc string) string {
	return string("\u001b[46m") + textobjc + string("\u001b[0m")
}

func BackWitRglr(textobjc string) string {
	return string("\u001b[47m") + textobjc + string("\u001b[0m")
}

// FOREGROUND + BOLD

func ForeBlkBold(textobjc string) string {
	return string("\u001b[30;1m") + textobjc + string("\u001b[0m")
}

func ForeRedBold(textobjc string) string {
	return string("\u001b[31;1m") + textobjc + string("\u001b[0m")
}

func ForeGrnBold(textobjc string) string {
	return string("\u001b[32;1m") + textobjc + string("\u001b[0m")
}

func ForeYlwBold(textobjc string) string {
	return string("\u001b[33;1m") + textobjc + string("\u001b[0m")
}

func ForeBluBold(textobjc string) string {
	return string("\u001b[34;1m") + textobjc + string("\u001b[0m")
}

func ForeMgtBold(textobjc string) string {
	return string("\u001b[35;1m") + textobjc + string("\u001b[0m")
}

func ForeCynBold(textobjc string) string {
	return string("\u001b[36;1m") + textobjc + string("\u001b[0m")
}

func ForeWitBold(textobjc string) string {
	return string("\u001b[37;1m") + textobjc + string("\u001b[0m")
}

// FOREGROUND + REGULAR

func ForeBlkRglr(textobjc string) string {
	return string("\u001b[30m") + textobjc + string("\u001b[0m")
}

func ForeRedRglr(textobjc string) string {
	return string("\u001b[31m") + textobjc + string("\u001b[0m")
}

func ForeGrnRglr(textobjc string) string {
	return string("\u001b[32m") + textobjc + string("\u001b[0m")
}

func ForeYlwRglr(textobjc string) string {
	return string("\u001b[33m") + textobjc + string("\u001b[0m")
}

func ForeBluRglr(textobjc string) string {
	return string("\u001b[34m") + textobjc + string("\u001b[0m")
}

func ForeMgtRglr(textobjc string) string {
	return string("\u001b[35m") + textobjc + string("\u001b[0m")
}

func ForeCynRglr(textobjc string) string {
	return string("\u001b[36m") + textobjc + string("\u001b[0m")
}

func ForeWitRglr(textobjc string) string {
	return string("\u001b[37m") + textobjc + string("\u001b[0m")
}

// CUSTOM STYLING

func ForeBold(textobjc string) string {
	return string("\u001b[1m") + textobjc + string("\u001b[0m")
}

func ForeUndl(textobjc string) string {
	return string("\u001b[4m") + textobjc + string("\u001b[0m")
}

func ForeInvr(textobjc string) string {
	return string("\u001b[7m") + textobjc + string("\u001b[0m")
}
