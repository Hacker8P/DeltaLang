section .data
        message db "Casiaush", 0xA
        lenmhgksak equ $ - message

section .bss
        f1sraxk resq 1
        f2sraxk resq 1
        f1posk resq 1
        f2posk resq 1
        cposk resq 1

section .text
        global _start

eqJmp:
        mov r10, [f1posk]

        mov rax, [f1sraxk]

        jmp r10

neqJmp:
        mov r10, [f2posk]

        mov rax, [f1sraxk]

        jmp r10

exit:
        mov rax, 60

        xor rdi, rdi
               
        syscall

scheduler:
        ;------------------------------------------
        ; SCHEDULER
        ;------------------------------------------

        mov r10, [cposk]

        xor r10, 1

        mov qword [cposk], r10

        cmp qword [cposk], 1

        je eqJmp

        jne neqJmp

y1:
        ;------------------------------------------
        ; YIELD FUNCTION 1
        ;------------------------------------------

        mov qword [f1sraxk], rax

        lea r10, [rel aft1]

        mov qword [f1posk], r10

        jmp post_yield

y2:
        ;------------------------------------------
        ; YIELD FUNCTION 2
        ;------------------------------------------

        mov qword [f2sraxk], rax

        lea r10, [rel aft2]

        mov qword [f2posk], r10

        jmp post_yield

yield:
        cmp r12, 1

        je y1

        jne y2

post_yield:
        jmp scheduler

f1:
        mov r12, 1
        mov rax, 1
        jmp yield

aft1:
        mov rdi, 1
        mov rsi, message
        mov rdx, lenmhgksak
        syscall

        jmp yield

f2:
        mov r12, 0
        mov rax, 1
        jmp yield

aft2:
        mov rdi, 1
        mov rsi, message
        mov rdx, lenmhgksak
        syscall

        jmp exit

_start:
        ;---------------------------------------------
        ; STARTING FUNCTION
        ;---------------------------------------------

        mov qword [cposk], 0

        lea r11, [rel f1]

        lea r10, [rel f2]

        mov qword [f1posk], r11

        mov qword [f2posk], r10

        jmp scheduler
