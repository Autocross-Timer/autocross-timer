terraform { 
  cloud { 
    
    organization = "MCO-Autocross" 

    workspaces { 
      name = "autocross-timer" 
    } 
  } 
}
