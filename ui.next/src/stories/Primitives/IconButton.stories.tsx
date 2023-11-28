import React from "react";
import { StoryObj, Meta } from "@storybook/react";
import { IconShare, IconSearch } from "@tabler/icons-react";

import { IconButton } from "./IconButton";

const meta: Meta<typeof IconButton> = {
	title: "Primitives/IconButton",
	component: IconButton,
	tags: ["autodocs"],
	argTypes: {
		icon: {
			options: ["Share", "Search"],
			mapping: {
				Share: <IconShare strokeWidth={1.25} size={20} />,
				Search: <IconSearch strokeWidth={1.25} size={20} />,
			},
		},
	},
};

export default meta;

type Story = StoryObj<typeof IconButton>;

export const Primary: Story = {
	args: {
		icon: <IconSearch strokeWidth={1.25} size={20} />,
		isPrimary: true,
	},
};

export const Secondary: Story = {
	args: {
		icon: <IconSearch strokeWidth={1.25} size={20} />,
		isPrimary: false,
	},
};
