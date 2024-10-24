---
title: Bluesky, hosting a PDS
date: 2024-10-23T10:47:37+02:00
draft: false
tags: []
---

As I mentioned in the **update** of [the last post](/posts/mastodon-bluesky-and-the-blue-bird/), I decided to give Bluesky a chance as long as I could host my own `PDS` to maintain the feeling of owning my data.

Setting aside the fact that I wanted to host the PDS within `Proxmox` and route it through an `nginx-proxy-manager`, it turned out to be easier than I thought. Initially, it didn’t work because the official stack includes not only the PDS but also `Caddy` and `Watchover`, causing conflicts between Caddy and `npm` (nginx-proxy-manager). I tried several approaches, but I wasn’t able to make it work.

### Install

Although I continued using a Proxmox CT, I assigned it a public IP directly so it wouldn’t route through NPM, and everything went smoothly:

```sh
$ apt install lsb-release
$ wget https://raw.githubusercontent.com/bluesky-social/pds/main/installer.sh
$ chmod +x installer.sh
$ ./installer.sh
* Detected supported distribution Debian 12
---------------------------------------
     Add DNS Record for Public IP
---------------------------------------
[...]
```

It’s also important to have a domain associated with that public IP via DNS. Although the instructions tell you to associate a wildcard, this is not necessary. I opted for `pds.oscarmlage.com` and simply created the corresponding DNS entry for that subdomain with the public IP:

```
pds.oscarmlage.com              A          51.77.72.104
```

The installer detects the IP itself, and after installing all the dependencies, it asks for the credentials of an account:

```
Detected public IP of this server: 51.77.72.104
Create a PDS user account? (y/N): y
Enter an email address (e.g. alice@pds.oscarmlage.com): oscarmlage@pds.oscarmlage.com
Enter a handle (e.g. alice.pds.oscarmlage.com): oscarmlage.pds.oscarmlage.com
Account created successfully!
-----------------------------
Handle   : oscarmlage.pds.oscarmlage.com
DID      : did:plc:XXXX
Password : XXXX
-----------------------------
Save this password, it will not be displayed again.
```

Everything should now be up and running. We can test it in several ways:

- Check health: `https://pds.oscarmlage.com/xrpc/_health`
- Check websocket: `wsdump.py "wss://pds.oscarmlage.com/xrpc/com.atproto.sync.subscribeRepos?cursor=0"`

We can now create more accounts and/or invitation codes from the console using the `pdsadmin` command:

```
# pdsadmin create-invite-code
pds-oscarmlage-com-xxxx-yyyy
```

### Account configuration

With the PDS now up and running, head over to [Bluesky's website](https://bsky.app/) to log in (if you already have credentials) or create a new account if you have an invite code:

{{< gallery match="gallery/create*" sortOrder="asc" rowHeight="150" margins="5" thumbnailResizeOptions="600x600 q90 Lanczos" previewType="blur" embedPreview="true" >}}


And what if you want to change the domain to something nicer? For example, I’d like my handle to be `@oscarmlage.com`instead of `@oscarmlage.pds.oscarmlage.com`. To do this, once you're logged in within bsky.app, go to *Settings* → *Change handle* → *I have my own domain* → Enter the new handle `@oscarmlage.com`, and it will show a *TXT* record that needs to be added to your *DNS* for ownership verification. If the validation is successful... voilà! You have your new custom handle!

{{< gallery match="gallery/change*" sortOrder="asc" rowHeight="150" margins="5" thumbnailResizeOptions="600x600 q90 Lanczos" previewType="blur" embedPreview="true" >}}

### Bridge with Mastodon

Another concern I have is managing duplicate posts. Although I’ll probably keep an eye on both platforms at first, I really don’t want to maintain two social networks, so the best option —for me— is to set up a bridge with Mastodon to automatically post the same content.

I’ve activated the bridge from Mastodon to Bluesky. To do this, I followed the bot `@bsky.brid.gy@bsky.brid.gy` from Mastodon. Once you start following it —from Mastodon— it automatically creates a Bluesky account under its domain with your Mastodon information, which becomes your bridge account (the bot will share the details in a *PM* with you). In my case, it's `@oscarmlage.mastodon.bofhers.es.ap.brid.gy`. And that’s it! Every time you post something on Mastodon, it will automatically be published to that Bluesky account.

{{< gallery match="gallery/bridge*" sortOrder="asc" rowHeight="150" margins="5" thumbnailResizeOptions="600x600 q90 Lanczos" previewType="blur" embedPreview="true" >}}

In the end, I’ve created two Bluesky accounts: one as my "main" account and another as the bridge with Mastodon. Feel free to follow me on either of them:

- [@oscarmlage.com](https://bsky.app/profile/oscarmlage.com)
- [@oscarmlage.mastodon.bofhers.es.ap.brid.gy](https://bsky.app/profile/oscarmlage.mastodon.bofhers.es.ap.brid.gy)

### Final thoughts

In [the last post](/posts/mastodon-bluesky-and-the-blue-bird/), I shared a story about a bird and a butterfly, with the elephant reminding us of the importance of staying grounded. But, as I reflect on it now, perhaps the butterflies, too, have something wonderful to offer. So, let's give Bluesky a chance to soar, without forgetting the solid ground that Mastodon has provided. After all, there's room for both beauty and stability in the skies we navigate.