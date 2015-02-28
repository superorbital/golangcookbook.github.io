---
layout: default
title: Go Cookbook
---

### Chapter 1:  Strings

* Introduction
* Accessing Substrings
* Establishing a Default Value
* Exchanging Values Without Using Temporary Variables
* Converting Between Characters and Values
* Using Named Unicode Characters
* Processing a String One Character at a Time
* Reversing a String by Word or Character
* Treating Unicode Combined Characters as Single Characters
* Canonicalizing Strings with Unicode Combined Characters
* Treating a Unicode String as Octets
* Expanding and Compressing Tabs
* Expanding Variables in User Input
* Controlling Case
* Properly Capitalizing a Title or Headline
* Interpolating Functions and Expressions Within Strings
* Indenting Here Documents
* Reformatting Paragraphs
* Escaping Characters
* Trimming Blanks from the Ends of a String
* Parsing Comma-Separated Data
* Constant Variables
* Soundex Matching
* Program: fixstyle
* Program: psgrep

### Chapter 2:  Numbers

* Introduction
* Checking Whether a String Is a Valid Number
* Rounding Floating-Point Numbers
* Comparing Floating-Point Numbers
* Operating on a Series of Integers
* Working with Roman Numerals
* Generating Random Numbers
* Generating Repeatable Random Number Sequences
* Making Numbers Even More Random
* Generating Biased Random Numbers
* Doing Trigonometry in Degrees, Not Radians
* Calculating More Trigonometric Functions
* Taking Logarithms
* Multiplying Matrices
* Using Complex Numbers
* Converting Binary, Octal, and Hexadecimal Numbers
* Putting Commas in Numbers
* Printing Correct Plurals
* Program: Calculating Prime Factors

### Chapter 3:  Dates and Times

* Introduction
* Finding Today's Date
* Converting DMYHMS to Epoch Seconds
* Converting Epoch Seconds to DMYHMS
* Adding to or Subtracting from a Date
* Difference of Two Dates
* Day in a Week/Month/Year or Week Number
* Parsing Dates and Times from Strings
* Printing a Date
* High-Resolution Timers
* Short Sleeps
* Program: hopdelta

### Chapter 4:  Arrays

* Introduction
* Specifying a List in Your Program
* Printing a List with Commas
* Changing Array Size
* Implementing a Sparse Array
* Iterating Over an Array
* Iterating Over an Array by Reference
* Extracting Unique Elements from a List
* Finding Elements in One Array but Not Another
* Computing Union, Intersection, or Difference of Unique Lists
* Appending One Array to Another
* Reversing an Array
* Processing Multiple Elements of an Array
* Finding the First List Element That Passes a Test
* Finding All Elements in an Array Matching Certain Criteria
* Sorting an Array Numerically
* Sorting a List by Computable Field
* Implementing a Circular List
* Randomizing an Array
* Program: words
* Program: permute

### Chapter 5:  Hashes

* Introduction
* Adding an Element to a Hash
* Testing for the Presence of a Key in a Hash
* Creating a Hash with Immutable Keys or Values
* Deleting from a Hash
* Traversing a Hash
* Printing a Hash
* Retrieving from a Hash in Insertion Order
* Hashes with Multiple Values per Key
* Inverting a Hash
* Sorting a Hash
* Merging Hashes
* Finding Common or Different Keys in Two Hashes
* Hashing References
* Presizing a Hash
* Finding the Most Common Anything
* Representing Relationships Between Data
* Program: dutree

### Chapter 6:  Pattern Matching

* Introduction
* Copying and Substituting Simultaneously
* Matching Letters
* Matching Words
* Commenting Regular Expressions
* Finding the Nth Occurrence of a Match
* Matching Within Multiple Lines
* Reading Records with a Separator
* Extracting a Range of Lines
* Matching Shell Globs as Regular Expressions
* Speeding Up Interpolated Matches
* Testing for a Valid Pattern
* Honoring Locale Settings in Regular Expressions
* Approximate Matching
* Matching from Where the Last Pattern Left Off
* Greedy and Non-Greedy Matches
* Detecting Doubled Words
* Matching Nested Patterns
* Expressing AND, OR, and NOT in a Single Pattern
* Matching a Valid Mail Address
* Matching Abbreviations
* Program: urlify
* Program: tcgrep
* Regular Expression Grab Bag

### Chapter 7:  File Access

* Introduction
* Opening a File
* Opening Files with Unusual Filenames
* Expanding Tildes in Filenames
* Making Perl Report Filenames in Error Messages
* Storing Filehandles into Variables
* Writing a Subroutine That Takes Filehandles as Built-ins Do
* Caching Open Output Filehandles
* Printing to Many Filehandles Simultaneously
* Opening and Closing File Descriptors by Number
* Copying Filehandles
* Creating Temporary Files
* Storing a File Inside Your Program Text
* Storing Multiple Files in the DATA Area
* Writing a Unix-Style Filter Program
* Modifying a File in Place with a Temporary File
* Modifying a File in Place with the -i Switch
* Modifying a File in Place Without a Temporary File
* Locking a File
* Flushing Output
* Doing Non-Blocking I/O
* Determining the Number of Unread Bytes
* Reading from Many Filehandles Without Blocking
* Reading an Entire Line Without Blocking
* Program: netlock
* Program: lockarea

### Chapter 8:  File Contents

* Introduction
* Reading Lines with Continuation Characters
* Counting Lines (or Paragraphs or Records) in a File
* Processing Every Word in a File
* Reading a File Backward by Line or Paragraph
* Trailing a Growing File
* Picking a Random Line from a File
* Randomizing All Lines
* Reading a Particular Line in a File
* Processing Variable-Length Text Fields
* Removing the Last Line of a File
* Processing Binary Files
* Using Random-Access I/O
* Updating a Random-Access File
* Reading a String from a Binary File
* Reading Fixed-Length Records
* Reading Configuration Files
* Testing a File for Trustworthiness
* Treating a File as an Array
* Setting the Default I/O Layers
* Reading or Writing Unicode from a Filehandle
* Converting Microsoft Text Files into Unicode
* Comparing the Contents of Two Files
* Pretending a String Is a File
* Program: tailwtmp
* Program: tctee
* Program: laston
* Program: Flat File Indexes

### Chapter 9:  Directories

* Introduction
* Getting and Setting Timestamps
* Deleting a File
* Copying or Moving a File
* Recognizing Two Names for the Same File
* Processing All Files in a Directory
* Globbing, or Getting a List of Filenames Matching a Pattern
* Processing All Files in a Directory Recursively
* Removing a Directory and Its Contents
* Renaming Files
* Splitting a Filename into Its Component Parts
* Working with Symbolic File Permissions Instead of Octal Values
* Program: symirror
* Program: lst

### Chapter 10:  Subroutines

* Introduction
* Accessing Subroutine Arguments
* Making Variables Private to a Function
* Creating Persistent Private Variables
* Determining Current Function Name
* Passing Arrays and Hashes by Reference
* Detecting Return Context
* Passing by Named Parameter
* Skipping Selected Return Values
* Returning More Than One Array or Hash
* Returning Failure
* Prototyping Functions
* Handling Exceptions
* Saving Global Values
* Redefining a Function
* Trapping Undefined Function Calls with AUTOLOAD
* Nesting Subroutines
* Writing a Switch Statement
* Program: Sorting Your Mail

### Chapter 11:  References and Records

* Introduction
* Taking References to Arrays
* Making Hashes of Arrays
* Taking References to Hashes
* Taking References to Functions
* Taking References to Scalars
* Creating Arrays of Scalar References
* Using Closures Instead of Objects
* Creating References to Methods
* Constructing Records
* Reading and Writing Hash Records to Text Files
* Printing Data Structures
* Copying Data Structures
* Storing Data Structures to Disk
* Transparently Persistent Data Structures
* Coping with Circular Data Structures Using Weak References
* Program: Outlines
* Program: Binary Trees

### Chapter 12:  Packages, Libraries, and Modules

* Introduction
* Defining a Module's Interface
* Trapping Errors in require or use
* Delaying use Until Runtime
* Making Variables Private to a Module
* Making Functions Private to a Module
* Determining the Caller's Package
* Automating Module Cleanup
* Keeping Your Own Module Directory
* Preparing a Module for Distribution
* Speeding Module Loading with SelfLoader
* Speeding Up Module Loading with Autoloader
* Overriding Built-in Functions
* Overriding a Built-in Function in All Packages
* Reporting Errors and Warnings Like Built-ins
* Customizing Warnings
* Referring to Packages Indirectly
* Using h2ph to Translate C #include Files
* Using h2xs to Make a Module with C Code
* Writing Extensions in C with Inline::C
* Documenting Your Module with Pod
* Building and Installing a CPAN Module
* Example: Module Template
* Program: Finding Versions and Descriptions of Installed Modules

### Chapter 14:  Database Access

* Introduction
* Making and Using a DBM File
* Emptying a DBM File
* Converting Between DBM Files
* Merging DBM Files
* Sorting Large DBM Files
* Storing Complex Data in a DBM File
* Persistent Data
* Saving Query Results to Excel or CSV
* Executing an SQL Command Using DBI
* Escaping Quotes
* Dealing with Database Errors
* Repeating Queries Efficiently
* Building Queries Programmatically
* Finding the Number of Rows Returned by a Query
* Using Transactions
* Viewing Data One Page at a Time
* Querying a CSV File with SQL
* Using SQL Without a Database Server
* Program: ggh—Grep Netscape Global History

### Chapter 15:  Interactivity

* Introduction
* Parsing Program Arguments
* Testing Whether a Program Is Running Interactively
* Clearing the Screen
* Determining Terminal or Window Size
* Changing Text Color
* Reading Single Characters from the Keyboard
* Ringing the Terminal Bell
* Using POSIX termios
* Checking for Waiting Input
* Reading Passwords
* Editing Input
* Managing the Screen
* Controlling Another Program with Expect
* Creating Menus with Tk
* Creating Dialog Boxes with Tk
* Responding to Tk Resize Events
* Removing the DOS Shell Window with Windows Perl/Tk
* Graphing Data
* Thumbnailing Images
* Adding Text to an Image
* Program: Small termcap Program
* Program: tkshufflepod
* Program: graphbox

### Chapter 16:  Process Management and Communication

* Introduction
* Gathering Output from a Program
* Running Another Program
* Replacing the Current Program with a Different One
* Reading or Writing to Another Program
* Filtering Your Own Output
* Preprocessing Input
* Reading STDERR from a Program
* Controlling Input and Output of Another Program
* Controlling the Input, Output, and Error of Another Program
* Communicating Between Related Processes
* Making a Process Look Like a File with Named Pipes
* Sharing Variables in Different Processes
* Listing Available Signals
* Sending a Signal
* Installing a Signal Handler
* Temporarily Overriding a Signal Handler
* Writing a Signal Handler
* Catching Ctrl-C
* Avoiding Zombie Processes
* Blocking Signals
* Timing Out an Operation
* Turning Signals into Fatal Errors
* Program: sigrand

### Chapter 17:  Sockets

* Introduction
* Writing a TCP Client
* Writing a TCP Server
* Communicating over TCP
* Setting Up a UDP Client
* Setting Up a UDP Server
* Using Unix Domain Sockets
* Identifying the Other End of a Socket
* Finding Your Own Name and Address
* Closing a Socket After Forking
* Writing Bidirectional Clients
* Forking Servers
* Pre-Forking Servers
* Non-Forking Servers
* Multitasking Server with Threads
* Writing a Multitasking Server with POE
* Writing a Multihomed Server
* Making a Daemon Server
* Restarting a Server on Demand
* Managing Multiple Streams of Input
* Program: backsniff
* Program: fwdport

### Chapter 18:  Internet Services

* Introduction
* Simple DNS Lookups
* Being an FTP Client
* Sending Mail
* Reading and Posting Usenet News Messages
* Reading Mail with POP3
* Simulating Telnet from a Program
* Pinging a Machine
* Accessing an LDAP Server
* Sending Attachments in Mail
* Extracting Attachments from Mail
* Writing an XML-RPC Server
* Writing an XML-RPC Client
* Writing a SOAP Server
* Writing a SOAP Client
* Program: rfrm
* Program: expn and vrfy

### Chapter 20:  Web Automation

* Introduction
* Fetching a URL from a Perl Script
* Automating Form Submission
* Extracting URLs
* Converting ASCII to HTML
* Converting HTML to ASCII
* Extracting or Removing HTML Tags
* Finding Stale Links
* Finding Fresh Links
* Using Templates to Generate HTML
* Mirroring Web Pages
* Creating a Robot
* Parsing a Web Server Log File
* Processing Server Logs
* Using Cookies
* Fetching Password-Protected Pages
* Fetching https:// Web Pages
* Resuming an HTTP GET
* Parsing HTML
* Extracting Table Data
* Program: htmlsub
* Program: hrefsub

