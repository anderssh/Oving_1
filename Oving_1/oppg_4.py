from threading import Thread;

i = 0;

def thread_function_1():

	global i;

	for j in range(0, 1000000):
		i = i + 1;

def thread_function_2():

	global i;
	
	for j in range(0, 1000000):
		i = i - 1;

def main():

	global i;

	thread_1 = Thread(target = thread_function_1, args = ());
	thread_2 = Thread(target = thread_function_2, args = ());

	thread_1.start();
	thread_2.start();

	thread_1.join();
	thread_2.join();

	print(i)

main();