# Hansel

Have you ever thought about friends and family while going about your day and wished they were there?
Perhaps you ate some delicious food that you wanted to share, or you simply saw a duck which reminded you of your friend.

With our busy and fast-paced lifestyle, it is difficult for us to meet up with friends, and the pandemic has only caused these catch up sessions to become even less frequent.
It is especially frustrating when you are even visiting the same places as your friends, but your paths do not cross because of conflicting schedules.
For example, with most lessons being carried out online, many students who do not stay on campus find that their friends do not turn up for classes on the same day as they do. 

Moreover, the immediacy of social media and messaging apps also mean that conversations are no longer something we actively anticipate or look forward to. 

Introducing... Hansel!

<p align="center">
   <img src="frontend/public/img/icons/android-chrome-192x192.png">
</p>

Hansel keeps you connected with your friends, closes the physical distance and helps to create shared memories even when you have missed each other at a place, all while maintaining a touch of novelty and surprise in your conversations.
Even if you find yourself going to campus alone, finding a message left for you by your friend there could brighten up your day.
The app name is a fond reference to Hansel, the eponymous character in Hansel and Gretel who leaves a trail of breadcrumbs to mark his path.
Similarly, our app allows users to leave a metaphorical trail of breadcrumbs for their friends that mark their path of travel. 

Users can leave media packages (‘iced gems’) for their friends when at a certain location, in the form of an image and an accompanying text.
After it is dropped, its location will be displayed on the receiver’s map and can only be opened if the receiver is in the vicinity of the gem.
After collecting and viewing the media package, the gem is then saved to be viewed again later if they wish.

Try Hansel [now](https://hansel-app.com/)!

## Team

| Member      | Role |
| ----------- | ----------- |
| Cai Jia Lin (A0205575A) | Full stack, UI, documentation |
| Ian Yong Yew Chuang (A0135451M) | Full stack, deployment, documentation |
| Sebastian Toh Shi Jian (A0196545R) | Full stack, deployment, documentation |
| Wu Weiming (A0210627N) | Frontend, UI, documentation |

## Getting Started

1. Install Protocol Buffer using the instructions [here](https://grpc.io/docs/protoc-installation/).
1. Follow the instructions in the respective frontend and backend `README`s.
1. Install Docker, Docker Compose, and run `docker-compose up` in the root directory to run an Envoy proxy.

   - If you are facing any difficulties getting the proxy to work, you may need replace `address: 127.0.0.1` with `address: host.docker.internal` or `address: docker.for.mac.localhost` in `envoy/envoy.yaml`
