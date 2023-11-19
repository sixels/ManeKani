import {
  ActionFunctionArgs,
  LoaderFunctionArgs,
  json,
  redirect,
} from '@remix-run/node';
import { useActionData, useLoaderData } from '@remix-run/react';
import classNames from 'classnames';
import {
  InvalidRequestError,
  UpdateUserDto,
  UsersAdapter,
} from 'manekani-core';
import { UsersDatabase } from 'manekani-infra-db';
import { TextInput } from '~/lib/components/form/Input';
import { Button } from '~/lib/components/general/Button';
import { database } from '~/lib/infra/db/db.server';
import { requireUserSession } from '~/lib/util/session';

export async function loader({ request }: LoaderFunctionArgs) {
  const { user } = await requireUserSession(request);

  if (user.isComplete) {
    throw redirect('/');
  }

  return json({ username: user.username, displayName: user.displayName });
}

export async function action({ request }: ActionFunctionArgs) {
  const { user } = await requireUserSession(request);
  if (user.isComplete) {
    throw redirect('/');
  }

  const formData = await request.formData();
  const fields = {
    username: formData.get('username')?.toString(),
    displayName: formData.get('display_name')?.toString(),
  };

  let errors: Record<string, string | null> = {
    username: fields.username ? null : 'Username is required',
    display_name: null,
    other: null,
  };

  if (fields.username) {
    const usersAdapter = new UsersAdapter(new UsersDatabase(database));
    const updateData: UpdateUserDto = {
      displayName: fields.displayName,
    };

    let isUserValid = false;
    if (fields.username == user.username) {
      isUserValid = true;
    } else {
      try {
        isUserValid = await usersAdapter.isUsernameAvailable(fields.username);
        updateData.username = fields.username;
      } catch (error) {
        console.error(error);
        if (error instanceof InvalidRequestError) {
          errors = {
            ...errors,
            ...handleValidationError(error),
          };
        }
      }
    }

    if (isUserValid) {
      try {
        console.log('updating user');
        await usersAdapter.updateUser(user.id, updateData);
      } catch (error) {
        console.error(error);
        if (error instanceof InvalidRequestError) {
          errors = {
            ...errors,
            ...handleValidationError(error),
          };
        } else {
          errors.other = "Couldn't complete the user profile";
        }
      }
    } else {
      errors.username =
        errors.username || `Username "${fields.username}" is not available`;
    }
  }

  const hasErrors = Object.values(errors).some((errorMessage) => errorMessage);
  if (hasErrors) {
    return json({ errors, fields });
  }

  return redirect('/');
}

export default function Component() {
  const { username, displayName } = useLoaderData<typeof loader>();

  const actionData = useActionData<typeof action>();

  // TODO: upload profile image
  // TODO: check if username is available everytime the user types

  return (
    <main className="flex max-w-screen-2xl md:p-2.5 gap-2.5 overflow-hidden h-screen mx-auto">
      <section className="left px-3.5 w-full">
        <h1 className="font-bold text-3xl text-neutral-900 leading-relaxed ">
          Complete Your Profile
        </h1>
        <p className="text-md text-neutral-800">
          You are almost there! We need a little more information from you.
        </p>
        <p className="text-md text-neutral-800">
          Don't worry, you can change them later.
        </p>
        <div className="mt-6">
          <form action="/complete-profile" method="post" className="space-y-6">
            <div>
              <label
                htmlFor="profile_picture"
                className="block text-sm font-medium text-neutral-500"
              >
                Profile Picture
              </label>
              <div className="flex flex-col items-center justify-center">
                <div className="w-36 h-36 rounded-full bg-neutral-200"></div>
                <button
                  type="button"
                  className="text-neutral-600 border transition-colors border-neutral-500 rounded-full px-4 py-1 focus:border-neutral-900  focus:text-neutral-900 focus:outline-none text-xs font-medium mt-3.5"
                >
                  Upload Picture
                </button>
              </div>
            </div>

            <TextInput
              id="display_name"
              name="display_name"
              label="Name"
              error={actionData?.errors.display_name}
              defaultValue={actionData?.fields.displayName ?? displayName}
            />

            <TextInput
              id="username"
              name="username"
              defaultValue={actionData?.fields.username ?? username}
              error={actionData?.errors.username}
              required
              label="Username"
            />

            {actionData?.errors.other && (
              <p className="mt-4 text-sm text-red-600">
                {actionData.errors.other}
              </p>
            )}

            <Button type="submit" isPrimary className="float-right">
              Continue
            </Button>
          </form>
        </div>
      </section>
      <div className="bg-neutral-200 h-full rounded-xl w-full max-w-xs lg:max-w-sm hidden md:block"></div>
    </main>
  );
}

function handleValidationError(error: InvalidRequestError) {
  console.error(JSON.stringify(error.context));
  const errors: Record<string, string | null> = {
    username: null,
    display_name: null,
  };

  if (error.context instanceof Object && 'errors' in error.context) {
    const {
      errors: validationErrors,
    }: {
      errors: {
        instancePath: string;
        schemaPath: string;
        keyword: string;
        params: { limit: number };
        message: string;
      }[];
    } = error.context as any;

    for (const ctx of validationErrors) {
      if (ctx.instancePath == '/displayName') {
        errors.display_name = `Name ${ctx.message}`;
      }
      if (ctx.instancePath == '/username' || ctx.instancePath == '') {
        errors.username = `Username ${ctx.message}`;
      }
    }
  }

  return errors;
}
