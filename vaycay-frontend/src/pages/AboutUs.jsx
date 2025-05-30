import { motion } from 'framer-motion';

const AboutUs = () => {
    return (
      <motion.div
        className="min-h-screen bg-gray-50"
        initial={{ opacity: 0, y: 20 }}
        animate={{ opacity: 1, y: 0 }}
        exit={{ opacity: 0, y: -20 }}
        transition={{ duration: 0.5 }}
      >
        {/* Hero Section */}
        <div className="py-12 pl-12 pr-4 sm:pl-16 sm:pr-6 lg:pl-24 lg:pr-8 bg-beige-600 text-left">
          <h1 className="text-4xl font-bold text-gray-900 mb-2">Your Hotel.</h1>
          <div className="flex items-center mb-6">
            <span className="text-4xl font-bold text-gray-900 mr-2">Your</span>
            <img
              src="/src/assets/Vacation..png"
              alt="Vacation"
              className="h-16 w-auto"
            />
          </div>
          <p className="text-xl text-gray-600  max-w-lg">
            We have over 240+ hotels waiting to give<br />
            you the best vacation ever.
          </p>
        </div>

        {/* What is Vacay Section */}
        <div className="relative py-12 pl-12 pr-4 sm:pl-16 sm:pr-6 lg:pl-24 lg:pr-8 text-left">
        {/* Semi-transparent background image only */}
        <div
            className="absolute inset-0 bg-cover bg-center z-0"
            style={{
            backgroundImage: `linear-gradient(to right, white, rgba(255, 255, 255, 0)), url('/src/assets/bgbg.jpg')`,
            }}
        ></div>

            <div className=" max-w-lg text-left relative z-10">
            <h2 className="text-3xl font-bold text-gray-900 mb-6">Our Story</h2>
            <p className="text-lg text-gray-600 mb-8">
                Vacay was founded in 2020 with a simple mission: to make hotel booking hassle-free and enjoyable. We started as a small team of travel enthusiasts who were frustrated with the complexity of planning accommodations for trips.
                <br /><br />
                Our platform was built to provide a straightforward, user-friendly experience that helps travelers find their perfect stay without the usual stress and complications. From budget-friendly options to luxury getaways, we curate only the best accommodations that meet our rigorous standards for quality and service.
                <br /><br />
                Today, Vacay has grown into a trusted platform helping thousands of travelers each month discover and book their ideal accommodations around the world. Our team has expanded, but our mission remains the same: to simplify travel planning and bring joy back to the journey.
            </p>
            </div>
        </div>

        {/* Stats Section */}
        <div className="py-16 px-4 sm:px-6 lg:px-8 bg-teal-600">
          <div className="max-w-4xl mx-auto grid grid-cols-2 md:grid-cols-4 gap-8 text-center">
            <div>
              <p className="text-4xl font-bold text-white">240+</p>
              <p className="text-white">Premium Hotels</p>
            </div>
            <div>
              <p className="text-4xl font-bold text-white">100+</p>
              <p className="text-white">Destinations</p>
            </div>
            <div>
              <p className="text-4xl font-bold text-white">15K+</p>
              <p className="text-white">Happy Customers</p>
            </div>
            <div>
              <p className="text-4xl font-bold text-white">4.8</p>
              <p className="text-white">Average Rating</p>
            </div>
          </div>
        </div>
  
        {/* Team Section */}
        <div className="py-16 px-4 sm:px-6 lg:px-8 bg-white">
          <div className="max-w-7xl mx-auto text-center">
            <h2 className="text-4xl font-bold text-gray-900 mb-4">Meet Our Team</h2>
            <p className="text-xl text-gray-600 mb-12 max-w-3xl mx-auto">
              The passionate people behind Vacay who work tirelessly to bring you the best hotel booking experience.
            </p>
            <div className="relative">
                <div className="flex space-x-8 pb-8 overflow-x-auto scrollbar-hide">
                    {[
                    {
                        name: "Jansen Choi Kai Xuan",
                        role: "Project Manager",
                        bio: "A results-driven leader with a sharp eye for detail, Jansen ensures smooth project execution and seamless coordination. He brings clarity, structure, and motivation to every phase of development.",
                        img: "/src/assets/team/jansen.jpg",
                    },
                    {
                        name: "Josh Edward Q. Lui",
                        role: "Backend Developer",
                        bio: "Josh specializes in building robust backend systems and APIs. Known for his logical approach and clean code, he ensures the performance and security of every application he works on.",
                        img: "/src/assets/team/josh.jpg",
                    },
                    {
                        name: "Jeskha Samantha Derama",
                        role: "Frontend Developer",
                        bio: "Jeskha translates designs into dynamic, responsive user interfaces. She is passionate about clean code, interactive design, and creating smooth digital experiences for end-users.",
                        img: "/src/assets/team/summi.jpg",
                    },
                    {
                        name: "Jose Rafael Achilles Delgado",
                        role: "UI/UX Developer",
                        bio: "With a keen sense of design and user psychology, Jose crafts intuitive and engaging interfaces. He focuses on creating seamless experiences that align with user needs and project goals.",
                        img: "/src/assets/team/jio.jpeg",
                    },
                    {
                        name: "Marie Louise Ty",
                        role: "UI/UX Developer",
                        bio: "Marie Louise brings creativity and functionality together through her user-centered designs. Marie focuses on accessibility, aesthetic coherence, and practical usability across all platforms.",
                        img: "/src/assets/team/maya.jpg",
                    },
                    ].map((member, index) => (
                        <div
                        key={index}
                        className="flex-shrink-0 w-80 bg-white rounded-lg shadow-md overflow-hidden"
                      >
                        <div className="h-48 bg-gray-200">
                          <img
                            src={member.img}
                            alt={member.name}
                            className="w-full h-full object-cover"
                          />
                        </div>
                        <div className="p-6">
                          <h3 className="text-xl font-bold text-gray-900">{member.name}</h3>
                          <p className="text-teal-600 mb-4">{member.role}</p>
                          <p className="text-gray-600">{member.bio}</p>
                        </div>
                      </div>
                    ))}
                </div>
            </div>
          </div>
        </div>
  
        {/* Values Section */}
        <div className="py-16 px-4 sm:px-6 lg:px-8 bg-beige-600">
          <div className="max-w-7xl mx-auto text-center">
            <h2 className="text-4xl font-bold text-teal-600 mb-4">Our Values</h2>
            <p className="text-xl text-gray-600 mb-12 max-w-3xl mx-auto">
              These core principles guide everything we do at Vacay.
            </p>
            <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
              {[
                {
                  title: "Trust & Transparency",
                  description:
                    "We believe in a living complexity that we are not our customers. No hidden fees, no misleading alternatives—just trust, information you can trust.",
                  img: "/src/assets/Check.png", // Path to check.png
                },
                {
                  title: "Customer-First Approach",
                  description:
                    "Every decision we make starts with one question: How does this benefit our customers? Your satisfaction is our top priority.",
                  img: "/src/assets/Arrow.png", // Path to arrow.png
                },
                {
                  title: "Passion for Travel",
                  description:
                    "We're travelers ourselves and understand what makes a stay memorable. This passion drives us to create exceptional experiences.",
                  img: "/src/assets/Play.png", // Path to play.png
                },
              ].map((value, index) => (
                <div
                  key={index}
                  className="bg-white h-72 flex flex-col justify-center items-center p-8 rounded-lg shadow-sm hover:shadow-md transition-shadow"
                >
                  {/* Image above the title */}
                  <div className="flex justify-center mb-4">
                    <img src={value.img} alt={value.title} className="h-12 w-12" />
                  </div>
                  <h3 className="text-xl font-bold text-gray-900 mb-4">{value.title}</h3>
                  <p className="text-gray-600 text-center">{value.description}</p>
                </div>
              ))}
            </div>
          </div>
        </div>
      </motion.div>
    );
  };
  
  export default AboutUs;