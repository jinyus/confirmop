 ConfirmOperation prompts the user to confirm the operation and  
 returns false if the user enters "n" or "no" and true if the user enters "y" or "yes".  
 If the user enters an invalid option and retry is true, the user will be prompted  
 until they enter a valid option. If retry is false, the function will return false.  

 desc is a short description of the operation. Can be empty.  
 confirmText is the question that the user should answer. Defaults to "confirm" if empty string is passed.  
 retry determines if the user should be prompted repeatedly until they enter a valid option.  

```
 example:

 userChoice := ConfirmOperation("Do you want to delete this file","proceed",true)

 // The user will see this prompt:
 //   Do you want to delete this file
 //   proceed? (y/no): 

 if userChoice{
   //safe to delete file
  }
```
