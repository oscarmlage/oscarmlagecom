---
title: "ZODB, how to properly delete a BLOB file"
date: 2014-11-25T22:48:23Z
draft: false
tags: [ "python", "code" ]
image: a25.jpg
---

<p>Due one of my&nbsp;<a href="http://www.pylonsproject.org/projects/pyramid/about">Pyramid</a>&nbsp;current projects, I was playing a bit with&nbsp;<a href="http://www.zodb.org/en/latest/">ZODB</a>&nbsp;and blob image storing. After reading here and there, I had no idea how to properly delete a filesystem stored image, blob field related. I though best way was to use&nbsp;<a href="http://nullege.com/codes/search/ZODB.blob.remove_committed">remove_committed()</a>&nbsp;<em>ZODB.blob</em>&nbsp;method with something like:</p>

```
def delete_avatar():
    image = myobj.avatar
    if isinstance(image, Blob):
        remove_committed(image.committed())
        myobj.avatar = None
```

<p>But after talking a bit with skilled&nbsp;<em>ZODB</em>&nbsp;people, they said that updating the database value and packing the database was enough to clear all the references in filesystem. First time ever I read something about "<em>packing</em>" a&nbsp;<em>ZODB</em>, so I had to probe it by myself.</p>

```
pshell$ profile.avatar
(ZODB.blob.Blob object at 0x108167230)
```

<p>Once we reached the image we can get it's path and some other interesting data for this test from&nbsp;<code>pshell</code>&nbsp;and shell:</p>

```
pshell$ profile.avatar.committed()
'/path/to/myproject/var/blobstorage/0x00/0x00/0x00/0x00/0x00/0x02/0xbd/0x68/0x03ab15928a1fb511.blob'
```


```
$ ls /path/to/myproject/var/blobstorage/0x00/0x00/0x00/0x00/0x00/0x02/0xbd/0x68/0x03ab15928a1fb511.blob
/0x03ab15928a1fb511.blob
$ find . -name *.blob | wc -l
12
$ du -s .
4784	.
```

<p>Nice!, if I open the&nbsp;<code>.blob</code>&nbsp;file with whatever image editor I get the proper image, next step is to enter in&nbsp;<code>pshell</code>&nbsp;mode again, update the database code to&nbsp;<code>None</code>,<em>commit</em>&nbsp;the transaction and see what's happening in the filesystem:</p>

```
pshell$ profile.avatar = None
pshell$ profile.avatar
pshell$ import transaction
pshell$ transaction.commit()
```


```
$ find . -name *.blob | wc -l
12
$ du -s .
4784	.
```

<p>Nothing seems to be happened, after the commit,&nbsp;<code>find</code>&nbsp;and&nbsp;<code>du</code>&nbsp;are showing the same information as before, last step is to&nbsp;<code>pack</code>&nbsp;the database...</p>

```
pshell$ from pyramid_zodbconn import get_connection
pshell$ conn = get_connection(request)
pshell$ conn.db().pack()
```


```
$ find . -name *.blob | wc -l
11
$ du -s .
4504	.
$ ls /path/to/myproject/var/blobstorage/0x00/0x00/0x00/0x00/0x00/0x02/0xbd/0x68/0x03ab15928a1fb511.blob
$
```

<p>That's it!, after a couple of seconds "<em>packing</em>", the filesystem seems to be updated, the file was properly deleted and all the information seems to be in a right state.</p>
<p>In the end seems that the&nbsp;<em>remove_commited()</em>&nbsp;function was not the right way to delete a&nbsp;<strong>BLOB</strong>&nbsp;file from&nbsp;<strong>ZODB</strong>, thumbs up for the packing!.</p>
