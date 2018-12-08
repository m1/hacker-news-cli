package mock

import (
	"encoding/json"
	"github.com/m1/hacker-news-cli/pkg/hn/posts"
)

var JsonString = `[{"id":18628295,"rank":1,"author":"wofo","comments":33,"points":102,"title":"Nginx on Wasmjit","posted":1544198430,"uri":"https://www.wasmjit.org/blog/nginx-on-wasmjit.html"},{"id":18626824,"rank":2,"author":"vanburen","comments":88,"points":197,"title":"AMD’s aggressive pricing update on the EPYC 7371","posted":1544184030,"uri":"https://www.servethehome.com/amd-epyc-7371-pricing-update-an-insane-value/"},{"id":18628140,"rank":3,"author":"deepaksurti","comments":105,"points":166,"title":"They Rejected Us","posted":1544194830,"uri":"https://rejected.us"},{"id":18628729,"rank":4,"author":"elithrar","comments":4,"points":18,"title":"Accidentally Moving from MacOS to Windows","posted":1544198670,"uri":"https://blog.questionable.services/article/accidentally-macos-wsl-windows-development/"},{"id":18628325,"rank":5,"author":"camtarn","comments":60,"points":52,"title":"The Trashing of Tesla Autopilot by Euro NCAP","posted":1544198430,"uri":"https://www.autoevolution.com/news/the-trashing-of-tesla-autopilot-by-euro-ncap-129683.html"},{"id":18626098,"rank":6,"author":"mbaytas","comments":142,"points":322,"title":"A blogging style guide","posted":1544176830,"uri":"https://robertheaton.com/2018/12/06/a-blogging-style-guide/"},{"id":18627020,"rank":7,"author":"se7entime","comments":63,"points":96,"title":"WordPress 5.0: A Gutenberg FAQ","posted":1544187630,"uri":"https://ma.tt/2018/11/a-gutenberg-faq/"},{"id":18626840,"rank":8,"author":"dnetesn","comments":47,"points":99,"title":"China prepares mission to land spacecraft on moon's far side","posted":1544184030,"uri":"https://phys.org/news/2018-12-china-mission-spacecraft-moon-side.html"},{"id":18627530,"rank":9,"author":"break_the_bank","comments":75,"points":127,"title":"Ask HN: How do you manage UI/UX for your side projects?","posted":1544191230,"uri":"https://news.ycombinator.com/item?id=18627530"},{"id":18626177,"rank":10,"author":"MikusR","comments":26,"points":100,"title":"Giving Mono Souper Powers","posted":1544176830,"uri":"https://www.mono-project.com/news/2018/12/06/souper/"},{"id":18623286,"rank":11,"author":"prismatic","comments":21,"points":61,"title":"Two images of the miners' strike, an instant apart: so which is the classic?","posted":1544180430,"uri":"https://www.theguardian.com/artanddesign/2018/dec/06/photographers-scrum-winning-image-picture"},{"id":18626281,"rank":12,"author":"tomduncalf","comments":21,"points":78,"title":"Show HN: Supportify – Support your top Spotify artists by buying on Bandcamp","posted":1544176830,"uri":"https://tomduncalf.github.io/supportify/"},{"id":18622644,"rank":13,"author":"chesterfield","comments":67,"points":99,"title":"Aleksandr Solzhenitsyn as he saw himself","posted":1544169630,"uri":"https://www.the-tls.co.uk/articles/public/solzhenitsyn-as-he-saw-himself/"},{"id":18625566,"rank":14,"author":"luu","comments":0,"points":15,"title":"Recurse Center: The return statement (2015)","posted":1544191230,"uri":"https://thewebivore.com/recurse-center-return-statement/"},{"id":18628852,"rank":15,"author":"syrusakbary","comments":2,"points":12,"title":"Show HN: The road of running Nginx with WebAssembly","posted":1544199390,"uri":"https://medium.com/@syrusakbary/running-nginx-with-webassembly-6353c02c08ac"},{"id":18627511,"rank":16,"author":"gadders","comments":33,"points":44,"title":"HCL buys Notes from IBM","posted":1544191230,"uri":"https://www.theregister.co.uk/2018/12/07/hcl_18bn_ibm_software/"},{"id":18625661,"rank":17,"author":"cryptoplot","comments":43,"points":82,"title":"Show HN: Coinmarketcap replacement with free JSON API","posted":1544166030,"uri":"http://cryptomarketplot.com"},{"id":18628059,"rank":18,"author":"Nuance","comments":0,"points":4,"title":"Stories of Steve Jobs in Safari Design Reviews","posted":1544194830,"uri":"https://donmelton.com/2014/04/10/memories-of-steve/"},{"id":18626132,"rank":19,"author":"DyslexicAtheist","comments":0,"points":19,"title":"The Protection of Information in Computer Systems (1974)","posted":1544184030,"uri":"https://www.cs.virginia.edu/~evans/cs551/saltzer/"},{"id":18625864,"rank":20,"author":"sohkamyung","comments":30,"points":44,"title":"DNS-Over-HTTPS (DoH) Operational and Privacy Issues","posted":1544169630,"uri":"https://www.ietf.org/blog/doh-operational-and-privacy-issues/"},{"id":18626316,"rank":21,"author":"jakub_g","comments":234,"points":465,"title":"Choose Firefox Now, or Later You Won't Get a Choice (2014)","posted":1544176830,"uri":"https://robert.ocallahan.org/2014/08/choose-firefox-now-or-later-you-wont.html"},{"id":18627035,"rank":22,"author":"kostspielig","comments":24,"points":61,"title":"Front-End Microservices","posted":1544187630,"uri":"https://jobs.zalando.com/tech/blog/front-end-micro-services/"},{"id":18622729,"rank":23,"author":"rozhok","comments":101,"points":421,"title":"Repl.it Multiplayer","posted":1544137230,"uri":"https://repl.it/site/blog/multi"},{"id":18612229,"rank":24,"author":"Hooke","comments":28,"points":70,"title":"Spooky Apollo: Apollo 8 and the CIA","posted":1544155230,"uri":"http://www.thespacereview.com/article/3617/1"},{"id":18620841,"rank":25,"author":"strmpnk","comments":5,"points":74,"title":"Why Systolic Architectures? (1982) [pdf]","posted":1544158830,"uri":"http://www.eecs.harvard.edu/~htk/publication/1982-kung-why-systolic-architecture.pdf"},{"id":18624630,"rank":26,"author":"cws","comments":23,"points":81,"title":"We busted a fake Chrome extension that was trying to steal data","posted":1544151630,"uri":"https://www.extrahop.com/company/blog/2018/fake-chrome-extension-threat-hunt/"},{"id":18618365,"rank":27,"author":"touristtam","comments":307,"points":736,"title":"Tesla's giant battery saved $40M during its first year, report says","posted":1544115630,"uri":"https://electrek.co/2018/12/06/tesla-battery-report/"},{"id":18629063,"rank":28,"author":"myinnerbanjo","comments":0,"points":3,"title":"Facial recognition: It’s time for action","posted":1544200650,"uri":"https://blogs.microsoft.com/on-the-issues/2018/12/06/facial-recognition-its-time-for-action/"},{"id":18627175,"rank":29,"author":"advert","comments":0,"points":0,"title":"Contract Simply (S17) Is Hiring Senior Engineers and a UI/UX Designer","posted":1544187630,"uri":"https://hire.withgoogle.com/public/jobs/contractsimplycom"},{"id":18618193,"rank":30,"author":"fanf2","comments":158,"points":710,"title":"A well-known URL for changing passwords","posted":1544115630,"uri":"https://github.com/WICG/change-password-url"},{"id":18628873,"rank":31,"author":"spatters","comments":0,"points":6,"title":"Standardizing on Keras","posted":1544199452,"uri":"https://medium.com/tensorflow/standardizing-on-keras-guidance-on-high-level-apis-in-tensorflow-2-0-bad2b04c819a"},{"id":18617791,"rank":32,"author":"svtrent","comments":277,"points":521,"title":"Lyft Files for IPO","posted":1544115632,"uri":"https://www.reuters.com/article/us-lyft-ipo/ride-hailing-firm-lyft-inc-files-for-ipo-idUSKBN1O51AA"},{"id":18626819,"rank":33,"author":"iluvdata","comments":5,"points":20,"title":"Ask HN: Data augmentation tools for 3d object detection","posted":1544184032,"uri":"https://news.ycombinator.com/item?id=18626819"},{"id":18626600,"rank":34,"author":"p_red","comments":0,"points":8,"title":"Show HN: ExifShot – A beautiful way to show metadata  of your photo","posted":1544180432,"uri":"https://exifshot.com/"},{"id":18618987,"rank":35,"author":"EvgeniyZh","comments":31,"points":364,"title":"StateOfTheArt.ai","posted":1544115632,"uri":"https://www.stateoftheart.ai/"},{"id":18620943,"rank":36,"author":"leafac","comments":0,"points":25,"title":"Understanding the Type of ‘call/cc’","posted":1544162432,"uri":"https://www.leafac.com/prose/understanding-the-type-of-call-cc/"},{"id":18624468,"rank":37,"author":"bubmiw","comments":63,"points":116,"title":"Six Detroit area doctors indicted in $500M health care fraud","posted":1544148032,"uri":"http://www.fox2detroit.com/news/local-news/six-detroit-area-doctors-indicted-in-500m-health-care-fraud"},{"id":18622062,"rank":38,"author":"igama","comments":56,"points":170,"title":"Kubernetes clusters being hijacked to mine cryptocurrencies","posted":1544130032,"uri":"https://blog.binaryedge.io/2018/12/06/kubernetes-being-hijacked-worldwide/"},{"id":18626922,"rank":39,"author":"bemmu","comments":45,"points":33,"title":"“Write Drunk, Edit Sober” Is Bad Advice","posted":1544187632,"uri":"https://goinswriter.com/write-drunk/"},{"id":18628114,"rank":40,"author":"kawera","comments":0,"points":6,"title":"FTC’s top consumer protection official can’t go after 120 large companies","posted":1544194832,"uri":"https://www.theverge.com/2018/12/6/18129572/facebook-uber-ftc-conflict-interest-andrew-smith"},{"id":18620978,"rank":41,"author":"Inufu","comments":70,"points":174,"title":"Show HN: AlphaZero Science paper","posted":1544126432,"uri":"https://deepmind.com/blog/alphazero-shedding-new-light-grand-games-chess-shogi-and-go/"},{"id":18623898,"rank":42,"author":"masonic","comments":144,"points":88,"title":"Fast Train to Failure: California’s Mismanaged High-Speed Rail Project","posted":1544148032,"uri":"https://www.city-journal.org/californias-high-speed-rail-project"},{"id":18624143,"rank":43,"author":"p1esk","comments":143,"points":150,"title":"LG Releases Gram 17 Laptop: An Ultra-Thin Notebook with a 17.3in Display","posted":1544144432,"uri":"https://www.anandtech.com/show/13681/lg-gram-17-available-ultra-thin-laptop-with-a-17-inch-display"},{"id":18619010,"rank":44,"author":"infodocket","comments":118,"points":198,"title":"Five-Year Trends Available for Median Income, Poverty and Internet Use","posted":1544115632,"uri":"https://www.census.gov/newsroom/press-releases/2018/2013-2017-acs-5year.html"},{"id":18620012,"rank":45,"author":"jorymackay","comments":34,"points":155,"title":"Why you need both rituals and routines to power your workday","posted":1544119232,"uri":"https://blog.rescuetime.com/workplace-routines-and-rituals/"},{"id":18616285,"rank":46,"author":"nickik","comments":22,"points":125,"title":"Keystone – Open-source Secure Hardware Enclave","posted":1544137232,"uri":"https://keystone-enclave.org"},{"id":18624870,"rank":47,"author":"arbuge","comments":69,"points":115,"title":"How Criminals Steal $37B a Year from America’s Elderly","posted":1544155232,"uri":"https://www.bloomberg.com/news/features/2018-05-03/america-s-elderly-are-losing-37-billion-a-year-to-fraud"},{"id":18624091,"rank":48,"author":"ArtWomb","comments":5,"points":62,"title":"Drawing a Picture Has a “Massive” Benefit for Memory versus Writing It Down","posted":1544144432,"uri":"https://digest.bps.org.uk/2018/11/22/the-act-of-drawing-something-has-a-massive-benefit-for-memory-compared-with-writing-it-down/"},{"id":18621166,"rank":49,"author":"pross356","comments":42,"points":123,"title":"Deepmind Produces a General-Purpose Games-Playing Machine","posted":1544126432,"uri":"https://spectrum.ieee.org/tech-talk/robotics/artificial-intelligence/mb"},{"id":18622516,"rank":50,"author":"__michaelg","comments":619,"points":1106,"title":"Goodbye, EdgeHTML","posted":1544133632,"uri":"https://blog.mozilla.org/blog/2018/12/06/goodbye-edge/"},{"id":18620309,"rank":51,"author":"harperlee","comments":53,"points":104,"title":"Better Clojure formatting","posted":1544133632,"uri":"http://tonsky.me/blog/clojurefmt"},{"id":18620404,"rank":52,"author":"Harj","comments":292,"points":440,"title":"The Rise of Microsoft Visual Studio Code","posted":1544122832,"uri":"https://triplebyte.com/blog/editor-report-the-rise-of-visual-studio-code"},{"id":18627273,"rank":53,"author":"alanwong","comments":79,"points":71,"title":"Huawei Reveals the Real Trade War with China","posted":1544191232,"uri":"https://www.bloomberg.com/news/articles/2018-12-06/huawei-bust-signals-the-real-u-s-trade-war-with-china"},{"id":18620879,"rank":54,"author":"stargrave","comments":134,"points":116,"title":"Why did Apple's Copland fail when so many 90s OS succeed, dominate world today","posted":1544122832,"uri":"https://liam-on-linux.livejournal.com/60604.html"},{"id":18617518,"rank":55,"author":"walterbell","comments":79,"points":199,"title":"Serial Port SDR","posted":1544115632,"uri":"https://hackaday.com/2018/12/06/your-usb-serial-adapter-just-became-a-sdr/"},{"id":18620668,"rank":56,"author":"JumpCrisscross","comments":0,"points":29,"title":"Arrest Shakes Huawei as Global Skepticism of Its Business Grows","posted":1544122832,"uri":"https://www.nytimes.com/2018/12/06/technology/huawei-arrest-meng-wanzhou.html"},{"id":18621513,"rank":57,"author":"arduinomancer","comments":167,"points":433,"title":"Facebook accused of striking 'secret deals over user data'","posted":1544126432,"uri":"https://www.bbc.com/news/technology-46456695"},{"id":18624128,"rank":58,"author":"Scarbutt","comments":197,"points":282,"title":"Neither PWA nor AMP are needed to make a website load fast","posted":1544144432,"uri":"http://tonsky.me/blog/pwa/"},{"id":18619754,"rank":59,"author":"jsheard","comments":164,"points":507,"title":"Rust 2018 is here, but what is it?","posted":1544119232,"uri":"https://hacks.mozilla.org/2018/12/rust-2018-is-here?"},{"id":18626897,"rank":60,"author":"dwighttk","comments":109,"points":103,"title":"The CRISPR Baby Scandal Gets Worse by the Day","posted":1544184032,"uri":"https://www.theatlantic.com/science/archive/2018/12/15-worrying-things-about-crispr-babies-scandal/577234/"}]`

func GetMockedPosts() posts.Posts {
	var ps posts.Posts
	postJs := JsonString
	_ = json.Unmarshal([]byte(postJs), &ps)

	return ps
}