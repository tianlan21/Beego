#!/usr/bin/env groovy

/*
  IBM Confidential OCO Source Materials
  (C) Copyright and Licensed by IBM Corp. 2017

  The source code for this program is not published or otherwise
  divested of its trade secrets  irrespective of what has
  been deposited with the U.S. Copyright Office.
*/


pipeline {
    agent any

    stages
    {
    	stage('Precheck') { 
		    steps { 
			    script {
                    echo 'Hello World 1'
                }
		    } 
	    }
        stage ('Invoke pipeline jobs') {
            steps {
            	script {
                    echo 'Hello World 2'
				
				}
            }
        }
    }
}
