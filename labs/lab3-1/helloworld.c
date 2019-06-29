/*
 * helloworld.c
 *
 *  Created on: Jun 22, 2019
 *      Author: elias
 */

#include <linux/module.h>

static int __init hello_init(void) {
	pr_info("Hello World init\n");
	return 0;
}

static int __exit hello_exit(void) {
	pr_info("Hello World exit\n");
	return 0;
}

module_init(hello_init);
module_exit(hello_exit);

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Elias Kopsiaftis <yemista@gmail.com>");
MODULE_DESCRIPTION("Lab 3-1");
