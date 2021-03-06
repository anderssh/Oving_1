// gcc -std=gnu99 -Wall -g -o oppg_4_c oppg_4.c -lpthread

#include <pthread.h>
#include <stdio.h>

int i = 0;

void* thread_function_1() {
	for (int j = 0; j < 1000000; j++) {
		i++;
	}
	return NULL;
}

void* thread_function_2() {
	for (int j = 0; j < 1000000; j++) {
		i--;
	}
	return NULL;
}

int main() {
	pthread_t thread_1;
	pthread_t thread_2;

	pthread_create(&thread_1, NULL, thread_function_1, NULL);
	pthread_create(&thread_2, NULL, thread_function_2, NULL);

	pthread_join(thread_1, NULL);
	pthread_join(thread_2, NULL);

	printf("%i\n", i);

	return 0;
}