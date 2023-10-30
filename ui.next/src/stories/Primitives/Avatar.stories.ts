import { Meta, StoryObj } from "@storybook/react";
import { Avatar } from "./Avatar";

const meta: Meta<typeof Avatar> = {
  title: "Primitives/Avatar",
  component: Avatar,
  tags: ["autodocs"],
};
export default meta;

type Story = StoryObj<typeof Avatar>;

export const Fallback: Story = {
  args: {
    isSquare: false,
    username: "sixels",
  },
};

export const CustomImage: Story = {
  args: {
    isSquare: false,
    username: "sixels",
    url: "https://i.pravatar.cc/150?img=9",
  },
};

export const Squared: Story = {
  args: {
    isSquare: true,
    username: "sixels",
    url: "https://i.pravatar.cc/150?img=9",
  },
};

export const Broken: Story = {
  args: {
    isSquare: false,
    username: "sixels",
    url: "http://example.com/not-found.png",
  },
};
