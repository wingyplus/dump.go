1# Dump.go
this project is fork repository from [googlecode](http://code.google.com/p/golang/)

*	dump, an utility to dump Go variables, similar to PHP's [var_dump](http://php.net/manual/en/function.var-dump.php)

### Example

	dump.Dump(&[][]int{[]int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}})
	
Outputs:

	&[][]int (len=3) {
	  []int (len=3) {
	    (int) 1,
	    (int) 2,
	    (int) 3
	  },
	  <[]int []int>,
	  <[]int []int>
	}

### Installation ###
Follow this command:

    $ make
    $ make install
