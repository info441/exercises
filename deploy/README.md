# deploy
Thus far, we've written and deployed simple web servers locally on our own machine. We've also
learned how to use `Docker` to containerize our applications. In this exercise, we will look to enable 
secure communications with HTTPS, containerize our web server and deploy it to the cloud. This exercise is meant
to be a demo. Students are encouraged to follow along and ask questions. 

## Prerequisite 
- You should have already signed up for an AWS account. (Sign up for AWS Educate to get ~$100 in credits)
If you haven’t done so already, please go ahead and [sign up](https://aws.amazon.com/). 
- You should have already registered SSH Keys. If you haven’t already done so, please refer to this [tutorial](https://drstearns.github.io/tutorials/deploy2aws/#secgeneratingandregisteringsshkeys). 
 
## Instructions 
You should've now completed all of the prerequisites. Follow the instructions below in order to complete the exercise.

1. Get a Domain Name. As a student, you get a free one from [Namecheap](https://nc.me/).   
2. Create AWS EC2 Instance. You should have already done this once in the [tutorial](https://drstearns.github.io/tutorials/deploy2aws/#secregisteryourpublickeywithawsec2) prior to coming to class. If you’ve already created one, feel free to use that for this exercise. If you haven’t done so, follow the instructions provided in the tutorial.
    - Please make sure you’ve installed Docker inside your instance. See https://drstearns.github.io/tutorials/deploy2aws/#secinstalldocker
    - Obtain certificates from Let's Encrypt. See https://drstearns.github.io/tutorials/https/#secrunletsencryptonamazonlinux2
    - Ensure that you’ve edited the security group in order to enable incoming request to port 80 (HTTP) and or 443 (HTTPS). See https://drstearns.github.io/tutorials/deploy2aws/#secedityoursecuritygroup
    - Create an Elastic IP and associate it with your EC2 instance. See https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html
3. Associate your Domain Name with your EC2 Instance. See https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-to-ec2-instance.html
    - Create Hosted Zone in Route 53
    - Create Record Set
        - Sub-domain (api, www)
    - Configure Nameservers on Namecheap
        - Copy the four nameserver entries in Route 53 and add to nameservers for Namecheap under custom DNS
        - This process takes some time for your IP to be registered 
4. 

