#include <stdint.h>
#include <stdlib.h>
#include <string.h>

//----- (0057DD90) --------------------------------------------------------
uint32_t* sub57DD90(uint32_t* this) {
	uint32_t* v1; // esi
	void* v2;     // eax

	v1 = this;
	v2 = calloc(1, 0x224u);
	*v1 = (uint32_t)v2;
	memset(v2, 0, 0x224u);
	return v1;
}

//----- (0057DDC0) --------------------------------------------------------
void sub57DDC0(void** this) { free(*this); }
