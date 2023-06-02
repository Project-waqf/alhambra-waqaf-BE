package email

import "os"

func GetBody(token string) string {

	// var url = "https://wakafalhambra.com?token=" + token
	email := os.Getenv("EMAIL_CALLBACK")
	var url =  email + "/new-password?token=" + token
	// var body = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd"><html xmlns="http://www.w3.org/1999/xhtml"><head><meta name="viewport" content="width=device-width, initial-scale=1.0" /><meta name="x-apple-disable-message-reformatting" /><meta http-equiv="Content-Type" content="text/html; charset=UTF-8" /><meta name="color-scheme" content="light dark" /><meta name="supported-color-schemes" content="light dark" /><title></title><style type="text/css" rel="stylesheet" media="all">/* Base ------------------------------ */@import url("https://fonts.googleapis.com/css?family=Nunito+Sans:400,700&display=swap");body {width: 100% !important;height: 100%;margin: 0;-webkit-text-size-adjust: none;}a {color: #3869D4;}a img {border: none;}td {word-break: break-word;}.preheader {display: none !important;visibility: hidden;mso-hide: all;font-size: 1px;line-height: 1px;max-height: 0;max-width: 0;opacity: 0;overflow: hidden;}/* Type ------------------------------ */body,td,th {font-family: "Nunito Sans", Helvetica, Arial, sans-serif;}h1 {margin-top: 0;color: #333333;font-size: 22px;font-weight: bold;text-align: left;}h2 {margin-top: 0;color: #333333;font-size: 16px;font-weight: bold;text-align: left;}h3 {margin-top: 0;color: #333333;font-size: 14px;font-weight: bold;text-align: left;}td,th {font-size: 16px;}p,ul,ol,blockquote {margin: .4em 0 1.1875em;font-size: 16px;line-height: 1.625;}p.sub {font-size: 13px;}/* Utilities ------------------------------ */.align-right {text-align: right;}.align-left {text-align: left;}.align-center {text-align: center;}.u-margin-bottom-none {margin-bottom: 0;}/* Buttons ------------------------------ */.button {background-color: #3869D4;border-top: 10px solid #3869D4;border-right: 18px solid #3869D4;border-bottom: 10px solid #3869D4;border-left: 18px solid #3869D4;display: inline-block;color: #FFF;text-decoration: none;border-radius: 3px;box-shadow: 0 2px 3px rgba(0, 0, 0, 0.16);-webkit-text-size-adjust: none;box-sizing: border-box;}.button--green {background-color: #22BC66;border-top: 10px solid #22BC66;border-right: 18px solid #22BC66;border-bottom: 10px solid #22BC66;border-left: 18px solid #22BC66;}.button--red {background-color: #FF6136;border-top: 10px solid #FF6136;border-right: 18px solid #FF6136;border-bottom: 10px solid #FF6136;border-left: 18px solid #FF6136;}@media only screen and (max-width: 500px) {.button {width: 100% !important;text-align: center !important;}}/* Attribute list ------------------------------ */.attributes {margin: 0 0 21px;}.attributes_content {background-color: #F4F4F7;padding: 16px;}.attributes_item {padding: 0;}/* Related Items ------------------------------ */.related {width: 100%;margin: 0;padding: 25px 0 0 0;-premailer-width: 100%;-premailer-cellpadding: 0;-premailer-cellspacing: 0;}.related_item {padding: 10px 0;color: #CBCCCF;font-size: 15px;line-height: 18px;}.related_item-title {display: block;margin: .5em 0 0;}.related_item-thumb {display: block;padding-bottom: 10px;}.related_heading {border-top: 1px solid #CBCCCF;text-align: center;padding: 25px 0 10px;}/* Discount Code ------------------------------ */.discount {width: 100%;margin: 0;padding: 24px;-premailer-width: 100%;-premailer-cellpadding: 0;-premailer-cellspacing: 0;background-color: #F4F4F7;border: 2px dashed #CBCCCF;}.discount_heading {text-align: center;}.discount_body {text-align: center;font-size: 15px;}/* Social Icons ------------------------------ */.social {width: auto;}.social td {padding: 0;width: auto;}.social_icon {height: 20px;margin: 0 8px 10px 8px;padding: 0;}/* Data table ------------------------------ */.purchase {width: 100%;margin: 0;padding: 35px 0;-premailer-width: 100%;-premailer-cellpadding: 0;-premailer-cellspacing: 0;}.purchase_content {width: 100%;margin: 0;padding: 25px 0 0 0;-premailer-width: 100%;-premailer-cellpadding: 0;-premailer-cellspacing: 0;}.purchase_item {padding: 10px 0;color: #51545E;font-size: 15px;line-height: 18px;}.purchase_heading {padding-bottom: 8px;border-bottom: 1px solid #EAEAEC;}.purchase_heading p {margin: 0;color: #85878E;font-size: 12px;}.purchase_footer {padding-top: 15px;border-top: 1px solid #EAEAEC;}.purchase_total {margin: 0;text-align: right;font-weight: bold;color: #333333;}.purchase_total--label {padding: 0 15px 0 0;}body {background-color: #F2F4F6;color: #51545E;}p {color: #51545E;}.email-wrapper {width: 100%;margin: 0;padding: 0;-premailer-width: 100%;-premailer-cellpadding: 0;-premailer-cellspacing: 0;background-color: #F2F4F6;}.email-content {width: 100%;margin: 0;padding: 0;-premailer-width: 100%;-premailer-cellpadding: 0;-premailer-cellspacing: 0;}/* Masthead ----------------------- */.email-masthead {padding: 25px 0;text-align: center;}.email-masthead_logo {width: 94px;}.email-masthead_name {font-size: 16px;font-weight: bold;color: #A8AAAF;text-decoration: none;text-shadow: 0 1px 0 white;}/* Body ------------------------------ */.email-body {width: 100%;margin: 0;padding: 0;-premailer-width: 100%;-premailer-cellpadding: 0;-premailer-cellspacing: 0;}.email-body_inner {width: 570px;margin: 0 auto;padding: 0;-premailer-width: 570px;-premailer-cellpadding: 0;-premailer-cellspacing: 0;background-color: #FFFFFF;}.email-footer {width: 570px;margin: 0 auto;padding: 0;-premailer-width: 570px;-premailer-cellpadding: 0;-premailer-cellspacing: 0;text-align: center;}.email-footer p {color: #A8AAAF;}.body-action {width: 100%;margin: 30px auto;padding: 0;-premailer-width: 100%;-premailer-cellpadding: 0;-premailer-cellspacing: 0;text-align: center;}.body-sub {margin-top: 25px;padding-top: 25px;border-top: 1px solid #EAEAEC;}.content-cell {padding: 45px;}/*Media Queries ------------------------------ */@media only screen and (max-width: 600px) {.email-body_inner,.email-footer {width: 100% !important;}}@media (prefers-color-scheme: dark) {body,.email-body,.email-body_inner,.email-content,.email-wrapper,.email-masthead,.email-footer {background-color: #333333 !important;color: #FFF !important;}p,ul,ol,blockquote,h1,h2,h3,span,.purchase_item {color: #FFF !important;}.attributes_content,.discount {background-color: #222 !important;}.email-masthead_name {text-shadow: none !important;}}:root {color-scheme: light dark;supported-color-schemes: light dark;}</style><!--[if mso]><style type="text/css">.f-fallback {font-family: Arial, sans-serif;}</style><![endif]--></head><body><span class="preheader">Use this link to reset your password. The link is only valid for 24 hours.</span><table class="email-wrapper" width="100%" cellpadding="0" cellspacing="0" role="presentation"><tr><td align="center"><table class="email-content" width="100%" cellpadding="0" cellspacing="0" role="presentation"><tr><td class="email-masthead"><a href="https://example.com" class="f-fallback email-masthead_name">[Product Name]</a></td></tr><!-- Email Body --><tr><td class="email-body" width="570" cellpadding="0" cellspacing="0"><table class="email-body_inner" align="center" width="570" cellpadding="0" cellspacing="0" role="presentation"><!-- Body content --><tr><td class="content-cell"><div class="f-fallback"><h1>Hi {{name}},</h1><p>You recently requested to reset your password for your [Product Name] account. Use the button below to reset it. <strong>This password reset is only valid for the next 24 hours.</strong></p><!-- Action --><table class="body-action" align="center" width="100%" cellpadding="0" cellspacing="0" role="presentation"><tr><td align="center"><!-- Border based buttonhttps://litmus.com/blog/a-guide-to-bulletproof-buttons-in-email-design --><table width="100%" border="0" cellspacing="0" cellpadding="0" role="presentation"><tr><td align="center"><a href="` + url + `" class="f-fallback button button--green" target="_blank">Reset your password</a></td></tr></table></td></tr></table><p>For security, this request was received from a {{operating_system}} device using {{browser_name}}. If you did not request a password reset, please ignore this email or <a href="{{support_url}}">contact support</a> if you have questions.</p><p>Thanks,<br>The [Product Name] team</p><!-- Sub copy --><table class="body-sub" role="presentation"><tr><td><p class="f-fallback sub">If you’re having trouble with the button above, copy and paste the URL below into your web browser.</p><p class="f-fallback sub">{{action_url}}</p></td></tr></table></div></td></tr></table></td></tr><tr><td><table class="email-footer" align="center" width="570" cellpadding="0" cellspacing="0" role="presentation"><tr><td class="content-cell" align="center"><p class="f-fallback sub align-center">[Company Name, LLC]<br>1234 Street Rd.<br>Suite 1234</p></td></tr></table></td></tr></table></td></tr></table></body></html>`
	var body = `
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <style>
        a{
            font-style: none;
            text-decoration: none;
        }

        .Lupa-Password {
        width: 800px;
        height: 250px;
        flex-grow: 0;
        padding: 40px 60px 40px;
        background-color: #fbfbfb;
        }
        .Hi-Admin {
        width: 147px;
        height: 40px;
        margin: 10px 20px 8px 0;
        margin-bottom: 20px;
        font-family: Inter;
        font-size: 32px;
        font-weight: 500;
        font-stretch: normal;
        font-style: normal;
        line-height: 1.25;
        letter-spacing: normal;
        text-align: left;
        color: #27303e;
        }
        .Line-15 {
        width: 700px;
        height: 1px;
        margin: 10px 0 10px;
        background-color: #c9c9c9;
        }
        .Permintaan-untuk-mengatur-ulang-password-telah-dibuat-Jika-Anda-tidak-membuat-permintaan-ini-abaik {
        width: 750px;
        height: 56px;
        margin: 48px 0 56px;
        font-family: Inter;
        font-size: 14px;
        font-weight: normal;
        font-stretch: normal;
        font-style: normal;
        line-height: 1.4;
        letter-spacing: normal;
        text-align: left;
        color: #27303e;
        }
        .Extra-Small-Label-Light-12px {
        width: 150px;
        height: 35px;
        flex-grow: 0;
        display: flex;
        margin-top: 10px;
        margin-bottom: 10px;
        flex-direction: row;
        justify-content: center;
        align-items: center;
        padding: 6px 30px;
        border-radius: 8px;
        background-color: #f98d3e;
        margin-left: 40%;
        border: none;
        }
        .Terima-kasih-Team-Al-Hambra {
        width: 1240px;
        height: 56px;
        margin: 48px 0 0;
        font-family: Inter;
        font-size: 14px;
        font-weight: normal;
        font-stretch: normal;
        font-style: normal;
        line-height: 1.4;
        letter-spacing: normal;
        text-align: left;
        color: #27303e;
        }
        .Large-Label-Heavy-18px {
        width: 138px;
        height: 27px;
        flex-grow: 0;
        font-family: Inter;
        font-size: 14px;
        margin-top: 1px;
        font-weight: 500;
        font-stretch: normal;
        font-style: normal;
        line-height: 1.5;
        letter-spacing: normal;
        text-align: center;
        color: #fbfbfb;
        }
        img.zyro-image-removebg-2 {
        width: 80px;
        height: 80px;
        margin: 0 0 0 450px;
        object-fit: contain;
        }
        .header{
            flex: auto;
            flex-direction: row;
        }
    </style>
</head>
<body>
    <div class="Lupa-Password">
        <div class="header">
            <span class="Hi-Admin">
                Hi, Admin
            </span>
            <!-- png -->
            <img src="https://ik.imagekit.io/alhambra/v1509_9269.png"
            srcset="img/zyro-image-removebg-2@2x.png 2x,
                    img/zyro-image-removebg-2@3x.png 3x"
            class="zyro-image-removebg-2">
        </div>
        <div class="Line-15"></div>
        <span class="Permintaan-untuk-mengatur-ulang-password-telah-dibuat-Jika-Anda-tidak-membuat-permintaan-ini-abaik">
            Permintaan untuk mengatur ulang password telah dibuat. Jika Anda tidak membuat permintaan ini, abaikan saja email ini. <br> Jika Anda memang membuat permintaan ini, harap setel ulang kata sandi Anda:
        </span>
		<button class="Extra-Small-Label-Light-12px">
        	<a href="` + url + `">
                <span class="Large-Label-Heavy-18px">
                    Reset Password
                </span>
			</a>
		</button>
        <span class="Terima-kasih-Team-Al-Hambra">
            Terima kasih,<br>
            Team Al Hambra
        </span>
    </div>
</body>
</html>
	`

	return body
}
