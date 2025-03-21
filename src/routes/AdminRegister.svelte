<script>
  import { onMount } from "svelte";
  import { push } from "svelte-spa-router";

  let formData = {
    username: "",
    email: "",
    password: "",
    confirmPassword: "",
  };
  let loading = false;
  let error = null;
  let validationErrors = {
    username: "",
    email: "",
    password: "",
    confirmPassword: "",
  };

  // Validate form fields
  function validateForm() {
    let isValid = true;

    // Reset validation errors
    validationErrors = {
      username: "",
      email: "",
      password: "",
      confirmPassword: "",
    };

    // Username validation
    if (!formData.username) {
      validationErrors.username = "Username is required";
      isValid = false;
    } else if (formData.username.length < 3) {
      validationErrors.username = "Username must be at least 3 characters";
      isValid = false;
    }

    // Email validation
    if (!formData.email) {
      validationErrors.email = "Email is required";
      isValid = false;
    } else {
      const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
      if (!emailRegex.test(formData.email)) {
        validationErrors.email = "Please enter a valid email address";
        isValid = false;
      }
    }

    // Password validation
    if (!formData.password) {
      validationErrors.password = "Password is required";
      isValid = false;
    } else if (formData.password.length < 6) {
      validationErrors.password = "Password must be at least 6 characters";
      isValid = false;
    }

    // Confirm password validation
    if (!formData.confirmPassword) {
      validationErrors.confirmPassword = "Please confirm your password";
      isValid = false;
    } else if (formData.password !== formData.confirmPassword) {
      validationErrors.confirmPassword = "Passwords do not match";
      isValid = false;
    }

    return isValid;
  }

  // Handle form submission
  async function handleSubmit(event) {
    event.preventDefault();

    if (!validateForm()) {
      return;
    }

    loading = true;
    error = null;

    try {
      // Prepare request data
      const requestData = {
        username: formData.username,
        email: formData.email,
        password: formData.password,
      };

      const response = await fetch("/api/auth/register", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(requestData),
      });

      const data = await response.json();

      if (!response.ok) {
        throw new Error(
          data.message || `Error ${response.status}: ${response.statusText}`
        );
      }

      // Store authentication token and user data
      sessionStorage.setItem("token", data.token);
      sessionStorage.setItem("user", JSON.stringify(data.user));

      // Navigate to store setup
      push("/admin/store/setup");
    } catch (err) {
      console.error("Registration failed:", err);
      error = err.message || "Registration failed. Please try again.";
      window.scrollTo(0, 0);
    } finally {
      loading = false;
    }
  }

  // Check if user is already logged in
  onMount(() => {
    const token =
      localStorage.getItem("token") || sessionStorage.getItem("token");
    if (token) {
      push("/admin/dashboard");
    }
  });
</script>

<svelte:head>
  <title>Admin Register | WhatsApp Catalogue</title>
</svelte:head>

<div class="min-h-screen bg-gray-50 flex flex-col">
  <!-- Header -->
  <header class="bg-white shadow">
    <div
      class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4 flex items-center justify-between"
    >
      <a class="flex items-center cursor-pointer" on:click={() => push("/")}>
        <div class="bg-[#25d366] p-2 rounded-full mr-3">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            stroke="white"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <path
              d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72 12.84 12.84 0 0 0 .7 2.81 2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45 12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"
            ></path>
          </svg>
        </div>
        <h1 class="text-xl font-semibold text-gray-800">WhatsApp Catalogue</h1>
      </a>
    </div>
  </header>

  <!-- Main Content -->
  <main class="flex-1 py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md mx-auto">
      <div class="text-center mb-8">
        <h2 class="text-3xl font-extrabold text-gray-900">
          Create your account
        </h2>
        <p class="mt-2 text-sm text-gray-600">
          Or
          <a
            href="/admin/login"
            class="font-medium text-[#25d366] hover:text-[#1da051]"
          >
            sign in to your existing account
          </a>
        </p>
      </div>

      {#if error}
        <div
          class="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded-lg text-sm mb-6"
        >
          {error}
        </div>
      {/if}

      <div class="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
        <form class="space-y-6" on:submit={handleSubmit}>
          <!-- Username field -->
          <div>
            <label
              for="username"
              class="block text-sm font-medium text-gray-700"
            >
              Username <span class="text-red-600">*</span>
            </label>
            <div class="mt-1">
              <input
                id="username"
                name="username"
                type="text"
                bind:value={formData.username}
                class="whatsapp-input"
                placeholder="your_username"
                required
              />
            </div>
            {#if validationErrors.username}
              <p class="mt-1 text-sm text-red-600">
                {validationErrors.username}
              </p>
            {/if}
          </div>

          <!-- Email field -->
          <div>
            <label for="email" class="block text-sm font-medium text-gray-700">
              Email <span class="text-red-600">*</span>
            </label>
            <div class="mt-1">
              <input
                id="email"
                name="email"
                type="email"
                bind:value={formData.email}
                class="whatsapp-input"
                placeholder="your@email.com"
                required
              />
            </div>
            {#if validationErrors.email}
              <p class="mt-1 text-sm text-red-600">{validationErrors.email}</p>
            {/if}
          </div>

          <!-- Password field -->
          <div>
            <label
              for="password"
              class="block text-sm font-medium text-gray-700"
            >
              Password <span class="text-red-600">*</span>
            </label>
            <div class="mt-1">
              <input
                id="password"
                name="password"
                type="password"
                bind:value={formData.password}
                class="whatsapp-input"
                placeholder="••••••••"
                required
              />
            </div>
            {#if validationErrors.password}
              <p class="mt-1 text-sm text-red-600">
                {validationErrors.password}
              </p>
            {:else}
              <p class="mt-1 text-xs text-gray-500">
                Password must be at least 6 characters
              </p>
            {/if}
          </div>

          <!-- Confirm Password field -->
          <div>
            <label
              for="confirmPassword"
              class="block text-sm font-medium text-gray-700"
            >
              Confirm Password <span class="text-red-600">*</span>
            </label>
            <div class="mt-1">
              <input
                id="confirmPassword"
                name="confirmPassword"
                type="password"
                bind:value={formData.confirmPassword}
                class="whatsapp-input"
                placeholder="••••••••"
                required
              />
            </div>
            {#if validationErrors.confirmPassword}
              <p class="mt-1 text-sm text-red-600">
                {validationErrors.confirmPassword}
              </p>
            {/if}
          </div>

          <div>
            <button
              type="submit"
              disabled={loading}
              class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-[#25d366] hover:bg-[#1da051] focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-[#25d366] disabled:opacity-50"
            >
              {#if loading}
                <svg
                  class="animate-spin -ml-1 mr-2 h-4 w-4 text-white"
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                >
                  <circle
                    class="opacity-25"
                    cx="12"
                    cy="12"
                    r="10"
                    stroke="currentColor"
                    stroke-width="4"
                  ></circle>
                  <path
                    class="opacity-75"
                    fill="currentColor"
                    d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
                  ></path>
                </svg>
                Creating account...
              {:else}
                Register
              {/if}
            </button>
          </div>
        </form>
      </div>

      <div class="mt-6 text-center">
        <p class="text-sm text-gray-600">
          By signing up, you agree to our
          <a href="#" class="font-medium text-[#25d366] hover:text-[#1da051]"
            >Terms of Service</a
          >
          and
          <a href="#" class="font-medium text-[#25d366] hover:text-[#1da051]"
            >Privacy Policy</a
          >.
        </p>
      </div>
    </div>
  </main>

  <!-- Footer -->
  <footer class="py-8 text-center text-gray-500 text-sm">
    <p>
      © {new Date().getFullYear()} WhatsApp Catalogue. All rights reserved.
    </p>
    <p class="mt-2">WhatsApp is a trademark of Meta Platforms, Inc.</p>
  </footer>
</div>
