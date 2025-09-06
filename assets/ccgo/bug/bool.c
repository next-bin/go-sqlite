#include <stdbool.h>
#include <assert.h>

// should transpile to something like "uint8 = BoolUint8(3 != 0)"
void literal() {
     bool a = 3;
     assert(a == 1);
}

// should transpile to: a = BoolUint8(a + 1)
void postInc() {
     bool a = true;
     a++;
     assert(a == 1);
}

// should transpile to: a = BoolUint8(a + 1)
void preInc() {
     bool a = true;
     ++a;
     assert(a == 1);

     --a;
     assert(a == 0);
}

// should apply BoolUint8
void postIncAssign() {
   bool a = true;
   bool x = a++;
	assert(x == 1 && a == 1);

	bool y = a--;
	assert(y == 1 && a == 0);
}

// should apply BoolUint8
void preIncAssign() {
   bool a = true;
   bool y = ++a;
	assert(y == 1 && a == 1);

	bool z = --a;
	assert(z == 0 && a == 0);
}

void incSizeOf() {
   bool a = true;
   assert(sizeof(a++) == 1);
	assert(sizeof(++a) == 1);
}

// should apply BoolUint8
void addAssign() {
   bool a = true;

   a += 1;
   assert(a == 1);

   // a+= 1 does the assignment & cast to bool first, then assigns true to x
	int x = a += 1;
	assert(x == 1 && a == 1);
}

void addAssign2() {
   bool a = true;

   // should cast to bool after applying +=
	int b = (a += 1) + (a += 1);
	assert(b == 2);
}

void addAssignSizeOf() {
   bool a = true;
   assert(sizeof(a += 1) == 1);
}

void integerPromotion() {
   bool a = true;
   assert(a + 1 == 2);
}

void wrapAround() {
   bool a = true;
   assert(a + 0xffffffff == 0);
	assert(!(bool)(a + 0xffffffff));
}

int main() {
   literal();
   postInc();
   preInc();
   postIncAssign();
   preIncAssign();
   incSizeOf();
   addAssign();
   addAssign2();
   addAssignSizeOf();
   integerPromotion();
   wrapAround();
}
