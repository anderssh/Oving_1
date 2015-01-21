// gcc -std=gnu99 -Wall -g -o oppg_1_c oppg_1.c -lpthread

#include <pthread.h>
#include <stdio.h>

int i = 0;

void* thread_function_1(void* mutex) {
	for (int j = 0; j < 1000000; j++) {
		pthread_mutex_lock(mutex);	
		i++;
		pthread_mutex_unlock(mutex);
	}
	return NULL;
}

void* thread_function_2(void* mutex) {
	for (int j = 0; j < 1000001; j++) {
		pthread_mutex_lock(mutex);			
		i--;
		pthread_mutex_unlock(mutex);
	}
	return NULL;
}

int main() {
	pthread_mutex_t mutex;
	pthread_mutex_init(&mutex, NULL);

	pthread_t thread_1;
	pthread_create(&thread_1, NULL, thread_function_1, &mutex);

	pthread_t thread_2;
	pthread_create(&thread_2, NULL, thread_function_2, &mutex);

	pthread_join(thread_1, NULL);
	pthread_join(thread_2, NULL);

	pthread_mutex_destroy(&mutex);

	printf("%i\n", i);

	return 0;
}