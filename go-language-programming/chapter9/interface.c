#include <stdio.h>
#include <stdlib.h>

typedef struct _TypeInfo {

} TypeInfo;

typedef struct _InterfaceInfo {

} InterfaceInfo;

typedef struct _IReadWriterTb1 {
    InterfaceInfo* inter;
    TypeInfo* type;
    int (*Read)(void* this, char* buf, int cb);
    int (*Writer)(void* this, char* buf, int cb);
} IReadWriterTb1;

typedef struct _IReadWriter {
    IReadWriterTb1* tab;
    void* data;
} IReadWriter;

InterfaceInfo g_InterfaceInfo_IReadWriter = {};

typedef struct _A {
    int a;
} A;

A* NewA(int params) {
    printf("NewA: %d\n", params);
    A* this = (A*)malloc(sizeof(A));
    this->a = params;
    return this;
}

int A_Read(A* this, char* buf, int cb) {
    printf("A_Read: %d\n", this->a);
    return cb;
}

int A_Write(A* this, char* buf, int cb) {
    printf("A_Write: %d\n", this->a);
    return cb;
}

TypeInfo g_TypeInfo_A = {};

typedef struct _B {
    A base;
} B;

B* NewB(int params) {
    printf("NewB: %d\n", params);
    B* this = (B*)malloc(sizeof(B));
    this->base.a = params;
    return this;
}

int B_Write(B* this, char* buf, int cb) {
    printf("B_Write: %d\n", this->base.a);
    return cb;
}

void B_Foo(B* this) {
    printf("B_Foo: ", this->base.a);
}

TypeInfo g_TypeInfo_B = {};

IReadWriterTb1 g_Itb1_IReadWriter_B = {
    &g_InterfaceInfo_IReadWriter,
    &g_TypeInfo_B,
    (int (*)(void* this, char* buf, int cb)) A_Read,
    (int (*)(void* this, char* buf, int cb)) B_Write
};

int main(void) {
    B* unnamed = NewB(8);
    //B_Write(unnamed, NULL, 10); 
    IReadWriter p = {
        &g_Itb1_IReadWriter_B,
        unnamed
    };
    p.tab->Read(p.data, NULL, 10);
    p.tab->Writer(p.data, NULL, 10);
    return 0;
}
