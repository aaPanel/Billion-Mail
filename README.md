<div align="center">
  <a name="readme-top"></a>
  <h1><a href="https://www.billionmail.com/" target="_blank">BillionMail 📧</a></h1>

## An Open-Source Mail Server, Email Marketing Solution for Smarter Campaigns

[![][license-shield]][license-link]
[![][docs-shield]][docs-link]
[![][github-release-shield]][github-release-link]
[![][github-stars-shield]][github-stars-link]

</div>
<br/>

## What is BillionMail?

BillionMail is a **future open-source Mail server, Email marketing platform** designed to help businesses and individuals manage their email campaigns with ease. Whether you're sending newsletters, promotional emails, or transactional messages, this tool will provide **full control** over your email marketing efforts. With features like **advanced analytics**, and **customer management**, you'll be able to create, send, and track emails like a pro.

![BillionMail Banner](https://www.billionmail.com/home.png?v1)

## How to use?
**Install Script:** (✅The script automatically installs all required runtime environments including Docker)
```shell
cd /opt
git clone https://github.com/aaPanel/BillionMail
cd BillionMail

# install with interact
bash install.sh
```


**Install with Docker:** (Please install Docker and docker-compose-plugin manually, and modify .env file)
```shell
# Navigate to /opt and clone the repository
cd /opt
git clone https://github.com/aaPanel/BillionMail

# Enter the BillionMail directory and copy the file
cd BillionMail
cp env_init .env

# Modify the default BILLIONMAIL_HOSTNAME value to your domain, e.g., mail.domain.com
# vi .env

# Initialize the certificate
mkdir ssl
cp -d -n ssl-self-signed/* ssl/

# Start BillionMail
docker compose up -d || docker-compose up -d

# Generate bm command
echo "$(pwd)" > /opt/PWD-Billion-Mail.txt
ln -sf $(pwd)/mail_users.sh /usr/bin/bm
chmod +x $(pwd)/mail_users.sh
```


<div align="center">
  <a href="https://www.youtube.com/watch?v=wGHfX1-7S_Y">
    <img src="https://img.youtube.com/vi/wGHfX1-7S_Y/maxresdefault.jpg" alt="" width="80%">
    <br />
    <img src="https://www.iconfinder.com/icons/317714/download/png/16" alt="YouTube" width="16"/>
    <b>Watch on YouTube</b>
  </a>
</div>


## Management script
Management help

`bm help`

View Login default info

`bm default`

Show domain DNS record

`bm show-record`

Update BillionMail
`bm update`


## WebMail

BillionMail has integrated **RoundCube**, you can access WebMail via `/roundcube/`.

## Why BillionMail?

Most email marketing platforms are either **expensive**, **closed-source**, or **lack essential features**. BillionMail aims to be different:

✅ **Fully Open-Source** – No hidden costs, no vendor lock-in.  
📊 **Advanced Analytics** – Track email delivery, open rates, click-through rates, and more.  
📧 **Unlimited Sending** – No restrictions on the number of emails you can send.  
🎨 **Customizable Templates** – Custom professional marketing templates for reuse.
🔒 **Privacy-First** – Your data stays with you, no third-party tracking.  
🚀 **Self-Hosted** – Run it on your own server for complete control.  

## How You Can Help 🌟

BillionMail is a **community-driven project**, and we need your support to get started! Here's how you can help:

1. **Star This Repository**: Show your interest by starring this repo.  
2. **Spread the Word**: Share BillionMail with your network—developers, marketers, and open-source enthusiasts.  
3. **Share Feedback**: Let us know what features you'd like to see in BillionMail by opening an issue or joining the discussion.  
4. **Contribute**: Once development begins, we'll welcome contributions from the community. Stay tuned for updates!

---

📧 **BillionMail – The Future of Open-Source Email Marketing.**

## Issues

If you encounter any issues or have feature requests, please [open an issue](https://github.com/aaPanel/BillionMail/issues). Be sure to include:

- A clear description of the problem or request.
- Steps to reproduce the issue (if applicable).
- Screenshots or error logs (if applicable).

## License

BillionMail is licensed under the **AGPLv3 License**. This means you can:

✅ Use the software for free.  
✅ Modify and distribute the code.  
✅ Use it privately without restrictions.

See the [LICENSE](LICENSE) file for more details.

---

📬 **BillionMail – Coming Soon. Star This Repo to Make It Happen Faster!**

<!-- BillionMail official link -->
[docs-link]: https://www.billionmail.com/

<!-- BillionMail Other link-->
[license-link]: https://www.gnu.org/licenses/gpl-3.0.html
[github-release-link]: https://github.com/aaPanel/BillionMail/releases/latest
[github-stars-link]: https://github.com/aaPanel/BillionMail
[github-issues-link]: https://github.com/aaPanel/BillionMail/issues

<!-- Shield link-->
[docs-shield]: https://img.shields.io/badge/documentation-148F76
[github-release-shield]: https://img.shields.io/github/v/release/aaPanel/BillionMail
[github-stars-shield]: https://img.shields.io/github/stars/aaPanel/BillionMail?color=%231890FF&style=flat-square   
[license-shield]: https://img.shields.io/github/license/aaPanel/BillionMail
